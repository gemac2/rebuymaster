package buyback

import (
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	orderID := c.Param("order_id")

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	buybacks := []models.Buyback{}
	if err := tx.Where("order_id = ?", orderID).All(&buybacks); err != nil {
		return errors.WithStack(errors.Wrap(err, "List - error getting all buybacks"))
	}

	filterBuybacks := models.FilterBuybacksByStopLoss(buybacks, order.OrderType)

	c.Set("buybacks", filterBuybacks)
	return c.Render(http.StatusOK, r.HTML("buybacks/index.plush.html"))
}
