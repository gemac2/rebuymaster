package order

import (
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Edit(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("order", order)
	return c.Render(http.StatusOK, r.HTML("orders/edit.plush.html"))
}
