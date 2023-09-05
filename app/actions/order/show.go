package order

import (
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func Show(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	buybacks := []models.Buyback{}
	if err := tx.Where("order_id = ?", orderID).All(&buybacks); err != nil {
		return errors.WithStack(errors.Wrap(err, "Show - error getting all buybacks"))
	}

	filterBuybacks := models.FilterBuybacksByStopLoss(buybacks, order.OrderType)

	c.Set("order", order)
	c.Set("buybacks", filterBuybacks)

	return c.Render(http.StatusOK, r.HTML("/orders/show.plush.html"))
}
