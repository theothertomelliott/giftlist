package actions

import (
	"github.com/gobuffalo/authrecipe/models"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to redirect to the starting page
func HomeHandler(c buffalo.Context) error {
	if uid := c.Session().Get("current_user_id"); uid == nil {
		u := models.User{}
		c.Set("user", u)
		return c.Render(200, r.HTML("auth/new.html"))
	}
	return c.Redirect(302, "/events")
}

// RoutesHandler is a handler to show all available routes
func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
