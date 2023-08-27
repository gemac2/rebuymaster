package order

import (
	"net/http"
	"rebuymaster/app/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Update(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(&order); err != nil {
		return err
	}

	order.UpdatedAt = time.Now()

	if err := tx.Update(&order); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Flash().Add("success", "Order updated successfully")

	return c.Redirect(http.StatusSeeOther, "/orders/%v/details", order.ID)

}
