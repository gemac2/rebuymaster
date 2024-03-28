package bybit

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
)

func Create(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("BYBIT_API_KEY")
	secretKey := os.Getenv("BYBIT_SECRET_KEY")

	client := bybit.NewClient().WithAuth(apiKey, secretKey)

	quantity := strconv.FormatFloat(order.CurrencyQuantity, 'f', -1, 64)
	price := strconv.FormatFloat(order.OrderPrice, 'f', 7, 64)

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

	if !order.IsBuybacksEnabled {
		slPrice := order.SetStopLossPrice()
		strSLPrice := strconv.FormatFloat(slPrice, 'f', -1, 64)
		orderParams.StopLoss = &strSLPrice
	}

	orderService := client.V5().Order()

	orderResp, err := orderService.CreateOrder(orderParams)
	if err != nil {
		fmt.Println("Error creating Bybit Order:", err)
		return err
	}

	fmt.Println("Bybit Order created successfully")
	fmt.Println("Order ID:", orderResp.Result.OrderID)

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/orders/%v/details", orderID))
	return c.Render(200, r.String("Bybit Order created"))
}
