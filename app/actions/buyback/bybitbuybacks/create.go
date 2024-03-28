package bybitbuybacks

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rebuymaster/app/models"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/hirokisan/bybit/v2"
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

	for i, fBuyback := range filterBuybacks {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		apiKey := os.Getenv("BYBIT_API_KEY")
		secretKey := os.Getenv("BYBIT_SECRET_KEY")

		client := bybit.NewClient().WithAuth(apiKey, secretKey)

		quantity := strconv.FormatFloat(fBuyback.CurrencyAmount, 'f', -1, 64)
		price := models.SetPriceForExchanges(order.CurrencyName, fBuyback.Price)

		bybitSymbol := models.GetBybitSymbol(order.CurrencyName)
		bybitSide := models.GetBybitSide(order.OrderType)
		positionMode := models.GetBybitPositionMode(order.OrderType)

		orderParams := bybit.V5CreateOrderParam{
			Category:    bybit.CategoryV5Linear,
			Symbol:      bybitSymbol,
			Side:        bybitSide,
			OrderType:   bybit.OrderTypeLimit,
			Qty:         quantity,
			Price:       &price,
			PositionIdx: positionMode,
		}

		if i == 0 {
			slPrice := models.SetPriceForExchanges(order.CurrencyName, filterBuybacks[len(filterBuybacks)-1].StopLossPrice)
			orderParams.StopLoss = &slPrice
		}

		orderService := client.V5().Order()

		orderResp, err := orderService.CreateOrder(orderParams)
		if err != nil {
			fmt.Println("Error creating Bybit Order:", err)
			return err
		}

		fmt.Println("Bybit Order created successfully")
		fmt.Println("Order ID:", orderResp.Result.OrderID)
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/buybacks/%v", orderID))
	return c.Render(200, r.String("Binance Order created"))
}
