package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to redirect to the starting page
func HomeHandler(c buffalo.Context) error {
	if uid := c.Session().Get("current_user_id"); uid == nil {
		return c.Redirect(301, "signin")
	}
	return c.Redirect(302, "/events")
}

// RoutesHandler is a handler to show all available routes
func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
