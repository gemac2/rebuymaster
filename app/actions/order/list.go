package order

import (
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	orders := []models.Order{}
	if err := tx.All(&orders); err != nil {
		return errors.WithStack(errors.Wrap(err, "List - error getting all orders"))
	}

	c.Set("orders", orders)

	return c.Render(http.StatusOK, r.HTML("orders/index.plush.html"))
}
