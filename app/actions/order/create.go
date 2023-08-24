package order

import (
	"net/http"
	"rebuymaster/app/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	order := models.Order{}
	if err := c.Bind(&order); err != nil {
		return err
	}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if err := tx.Create(&order); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Flash().Add("success", "Order created successfully")

	return c.Redirect(http.StatusSeeOther, "/orders")
}
