package binance

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
)

func SetOrder(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	binanceSide := models.GetBinanceSide(order.OrderType)

	quantity := strconv.FormatFloat(order.CurrencyQuantity, 'f', -1, 64)
	price := models.SetPriceForExchanges(order.CurrencyName, order.OrderPrice)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	client := futures.NewClient(apiKey, secretKey)

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

	if !order.IsBuybacksEnabled {
		slPrice := order.SetStopLossPrice()
		strSLPrice := models.SetPriceForExchanges(order.CurrencyName, slPrice)
		binanceSLSide := models.GetBinanceStopLossSide(order.OrderType)
		stopOrder := client.NewCreateOrderService().
			Symbol(order.CurrencyName).
			Side(binanceSLSide).
			Type(futures.OrderTypeStopMarket).
			Quantity(quantity).
			StopPrice(strSLPrice)

		respStop, err := stopOrder.Do(c)
		if err != nil {
			log.Println("Error creating stop loss order:", err)
			os.Exit(1)
		}

		log.Printf("Stop Loss Futures Order ID: %d\n", respStop.OrderID)
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/orders/%v/details", orderID))
	return c.Render(200, r.String("Binance Order created"))
}
