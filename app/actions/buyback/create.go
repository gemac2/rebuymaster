package buyback

import (
	"fmt"
	"math"
	"net/http"
	"rebuymaster/app/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	orderID := c.Param("order_id")

	order := models.Order{}
	if err := tx.Find(&order, orderID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	buybacks := models.GenerateBuybacks(order)

	for i, buyback := range buybacks {
		if i == 0 {
			continue
		}

		buyback.Price = math.Round(buyback.Price*1e7) / 1e7

		buyback.OrderID = order.ID
		buyback.CreatedAt = time.Now()
		buyback.UpdatedAt = time.Now()

		if err := tx.Create(&buyback); err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
	}

	c.Flash().Add("success", "Buybacks created successfully")
	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/orders/%v/details", order.ID))

	return c.Render(200, r.String("Buybacks created"))
}
