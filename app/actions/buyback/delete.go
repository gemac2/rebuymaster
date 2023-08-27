package buyback

import (
	"fmt"
	"net/http"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func DeleteBuybacks(c buffalo.Context) error {
	orderID := c.Param("order_id")
	tx := c.Value("tx").(*pop.Connection)

	buybacks := []models.Buyback{}
	if err := tx.Where("order_id = ?", orderID).All(&buybacks); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	for _, buyback := range buybacks {
		if err := tx.Destroy(&buyback); err != nil {
			return errors.WithStack(errors.Wrap(err, "DeleteBuyback - Error deleting buyback"))
		}
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/orders/%v/details", orderID))
	return c.Render(200, r.String("Buybacks were delete successfully"))
}
