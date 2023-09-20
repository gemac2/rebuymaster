package order

import (
	"net/http"
	"rebuymaster/app/actions/paginator"
	"rebuymaster/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	q := tx.PaginateFromParams(c.Params())

	if err := paginator.ApplyLimitAndOffset(q, c); err != nil {
		return c.Error(http.StatusInternalServerError, errors.Wrapf(err, "could not apply the limit and offset to the given query"))
	}

	orders := []models.Order{}

	q = q.Order("created_at ASC")

	if err := q.All(&orders); err != nil {
		return errors.WithStack(errors.Wrap(err, "List - error while getting all orders from db"))
	}

	paginator := paginator.NewPagination(int32(q.Paginator.PerPage), int32(q.Paginator.Offset), int32(len(orders)), int64(q.Paginator.TotalEntriesSize))

	c.Set("pagination", paginator)
	c.Set("orders", orders)

	return c.Render(http.StatusOK, r.HTML("orders/index.plush.html"))
}
