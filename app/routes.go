package app

import (
	"net/http"

	"rebuymaster/app/actions/buyback"
	"rebuymaster/app/actions/buyback/binancebuybacks"
	"rebuymaster/app/actions/buyback/bybitbuybacks"
	"rebuymaster/app/actions/order"
	"rebuymaster/app/actions/order/binance"
	"rebuymaster/app/actions/order/bybit"
	"rebuymaster/app/actions/order/graphics"
	"rebuymaster/app/middleware"
	"rebuymaster/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(app *buffalo.App) {
	app.Use(middleware.RequestID)
	app.Use(middleware.Database)
	app.Use(middleware.ParameterLogger)
	// app.Use(middleware.CSRF)

	app.GET("/", order.List)

	orders := app.Group("/orders")
	orders.GET("/", order.List)
	orders.GET("/new", order.New)
	orders.GET("/{order_id:[-0-9a-z]+}/details", order.Show)
	orders.GET("/{order_id:[-0-9a-z]+}/edit", order.Edit)
	orders.PUT("/{order_id:[-0-9a-z]+}/update", order.Update)
	orders.POST("/create", order.Create)
	orders.DELETE("/{order_id:[-0-9a-z]+}/delete", order.DeleteOrder)
	orders.GET("/graphics", graphics.GenerateGraphics)
	// route to create order in bybit
	orders.POST("/{order_id:[-0-9a-z]+}/set-order", bybit.Create)

	// route to create binance order
	orders.POST("/{order_id:[-0-9a-z]+}/create-order", binance.SetOrder)

	buybacks := app.Group("/buybacks")
	buybacks.GET("/{order_id:[-0-9a-z]+}", buyback.List)
	buybacks.POST("/{order_id:[-0-9a-z]+}/create", buyback.Create)
	buybacks.DELETE("/{order_id:[-0-9a-z]+}/delete", buyback.DeleteBuybacks)

	// route to create binance buybacks
	buybacks.POST("/{order_id:[-0-9a-z]+}/binance/create", binancebuybacks.CreateBuybacks)

	// route to create bybit buybacks
	buybacks.POST("/{order_id:[-0-9a-z]+}/bybit/create", bybitbuybacks.CreateBuybacks)

	app.ServeFiles("/", http.FS(public.FS()))
}
