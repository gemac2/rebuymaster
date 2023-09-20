package render

import (
	"fmt"
	"net/url"
	"rebuymaster/app/templates"
	"rebuymaster/public"

	"github.com/gobuffalo/buffalo/render"
	"github.com/wawandco/ox/pkg/buffalotools"
)

// Engine for rendering across the app, it provides
// the base for rendering HTML, JSON, XML and other formats
// while also defining thing like the base layout.
var Engine = render.New(render.Options{
	HTMLLayout:  "application.plush.html",
	TemplatesFS: templates.FS(),
	AssetsFS:    public.FS(),
	Helpers:     Helpers,
})

// Helpers available for the plush templates, there are
// some helpers that are injected by Buffalo but this is
// the list of custom Helpers.
var Helpers = map[string]interface{}{
	// partialFeeder is the helper used by the render engine
	// to find the partials that will be used, this is important
	"partialFeeder": buffalotools.NewPartialFeeder(templates.FS()),
	"linkWith":      linkWith,
	"linkWithout":   linkWithout,
}

func linkWith(current string, fields map[string]interface{}) string {
	path, err := url.Parse(current)
	if err != nil {
		return current
	}

	q := path.Query()
	for field, val := range fields {
		q.Set(field, fmt.Sprintf("%v", val))
	}

	path.RawQuery = q.Encode()

	return path.String()
}

func linkWithout(current string, fields ...string) string {
	path, err := url.Parse(current)
	if err != nil {
		return current
	}

	q := path.Query()
	for _, field := range fields {
		q.Del(field)
	}
	path.RawQuery = q.Encode()

	return path.String()
}
