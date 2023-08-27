package order

import (
	"fmt"
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func DeleteOrder(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(&order); err != nil {
		return errors.WithStack(errors.Wrap(err, "DeleteOrder - Error deleting order"))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/orders"))
	return c.Render(200, r.String("The Order was delete"))
}
