package buyback

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func List(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("buybacks/index.plush.html"))
}
