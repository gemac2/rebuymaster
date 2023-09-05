package binancebuybacks

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rebuymaster/app/models"
	"strconv"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func CreateBuybacks(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	buybacks := []models.Buyback{}
	if err := tx.Where("order_id = ?", orderID).All(&buybacks); err != nil {
		return errors.WithStack(errors.Wrap(err, "List - error getting all buybacks"))
	}

	filterBuybacks := models.FilterBuybacksByStopLoss(buybacks, order.OrderType)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	client := futures.NewClient(apiKey, secretKey)

	for _, fBuyback := range filterBuybacks {
		binanceSide := models.GetBinanceSide(order.OrderType)
		quantity := strconv.FormatFloat(fBuyback.CurrencyAmount, 'f', -1, 64)
		price := models.SetPriceForExchanges(order.CurrencyName, fBuyback.Price)

		binanceOrder := client.NewCreateOrderService().
			Symbol(order.CurrencyName).
			Side(binanceSide).
			Type(futures.OrderTypeLimit).
			Quantity(quantity).
			Price(price).
			TimeInForce(futures.TimeInForceTypeGTC)

		resp, err := binanceOrder.Do(c)
		if err != nil {
			log.Println("Error creating futures order:", err)
			return c.Error(http.StatusInternalServerError, err)
		}

		log.Printf("Futures Order ID: %d\n", resp.OrderID)
	}

	binanceSLSide := models.GetBinanceStopLossSide(order.OrderType)
	slPrice := models.SetPriceForExchanges(order.CurrencyName, filterBuybacks[len(filterBuybacks)-1].StopLossPrice)
	quantity := strconv.FormatFloat((filterBuybacks[len(filterBuybacks)-1].CurrencyAmount * 2), 'f', -1, 64)
	stopOrder := client.NewCreateOrderService().
		Symbol(order.CurrencyName).
		Side(binanceSLSide).
		Type(futures.OrderTypeStopMarket).
		Quantity(quantity).
		StopPrice(slPrice)

	respStop, err := stopOrder.Do(c)
	if err != nil {
		log.Println("Error creating stop loss order:", err)
		os.Exit(1)
	}

	log.Printf("Stop Loss Futures Order ID: %d\n", respStop.OrderID)

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/buybacks/%v", orderID))
	return c.Render(200, r.String("Binance Order created"))
}
