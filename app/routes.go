package app

import (
	"net/http"

	"rebuymaster/app/actions/buyback"
	"rebuymaster/app/actions/order"
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
	orders.GET("/{order_id:[-0-9a-z]+}/details/", order.Show)
	orders.POST("/create", order.Create)

	buybacks := app.Group("/buybacks")
	buybacks.GET("/{order_id:[-0-9a-z]+}", buyback.List)
	buybacks.POST("/{order_id:[-0-9a-z]+}/create", buyback.Create)

	app.ServeFiles("/", http.FS(public.FS()))
}
