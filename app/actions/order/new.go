package order

import (
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
)

func New(c buffalo.Context) error {
	order := models.Order{}

	c.Set("order", order)
	return c.Render(http.StatusOK, r.HTML("orders/new.plush.html"))
}
