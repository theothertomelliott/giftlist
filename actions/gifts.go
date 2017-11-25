package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	"github.com/theothertomelliott/giftlist/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Gift)
// DB Table: Plural (gifts)
// Resource: Plural (Gifts)
// Path: Plural (/gifts)
// View Template Folder: Plural (/templates/gifts/)

// GiftsResource is the resource for the Gift model
type GiftsResource struct {
	buffalo.Resource
}

// List gets all Gifts. This function is mapped to the path
// GET /gifts
func (v GiftsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	gifts := &models.Gifts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Gifts from the DB
	if err := q.All(gifts); err != nil {
		return errors.WithStack(err)
	}

	// Make Gifts available inside the html template
	c.Set("gifts", gifts)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("gifts/index.html"))
}

// Show gets the data for one Gift. This function is mapped to
// the path GET /gifts/{gift_id}
func (v GiftsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Gift
	gift := &models.Gift{}

	// To find the Gift the parameter gift_id is used.
	if err := tx.Find(gift, c.Param("gift_id")); err != nil {
		return c.Error(404, err)
	}

	// Make gift available inside the html template
	c.Set("gift", gift)

	return c.Render(200, r.HTML("gifts/show.html"))
}

// New renders the form for creating a new Gift.
// This function is mapped to the path GET /gifts/new
func (v GiftsResource) New(c buffalo.Context) error {
	// Make gift available inside the html template
	c.Set("gift", &models.Gift{})

	return c.Render(200, r.HTML("gifts/new.html"))
}

// Create adds a Gift to the DB. This function is mapped to the
// path POST /gifts
func (v GiftsResource) Create(c buffalo.Context) error {
	// Allocate an empty Gift
	gift := &models.Gift{}

	// Bind gift to the html form elements
	if err := c.Bind(gift); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(gift)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make gift available inside the html template
		c.Set("gift", gift)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("gifts/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Gift was created successfully")

	// and redirect to the gifts index page
	return c.Redirect(302, "/gifts/%s", gift.ID)
}

// Edit renders a edit form for a Gift. This function is
// mapped to the path GET /gifts/{gift_id}/edit
func (v GiftsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Gift
	gift := &models.Gift{}

	if err := tx.Find(gift, c.Param("gift_id")); err != nil {
		return c.Error(404, err)
	}

	// Make gift available inside the html template
	c.Set("gift", gift)
	return c.Render(200, r.HTML("gifts/edit.html"))
}

// Update changes a Gift in the DB. This function is mapped to
// the path PUT /gifts/{gift_id}
func (v GiftsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Gift
	gift := &models.Gift{}

	if err := tx.Find(gift, c.Param("gift_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Gift to the html form elements
	if err := c.Bind(gift); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(gift)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make gift available inside the html template
		c.Set("gift", gift)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("gifts/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Gift was updated successfully")

	// and redirect to the gifts index page
	return c.Redirect(302, "/gifts/%s", gift.ID)
}

// Destroy deletes a Gift from the DB. This function is mapped
// to the path DELETE /gifts/{gift_id}
func (v GiftsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Gift
	gift := &models.Gift{}

	// To find the Gift the parameter gift_id is used.
	if err := tx.Find(gift, c.Param("gift_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(gift); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Gift was destroyed successfully")

	// Redirect to the gifts index page
	return c.Redirect(302, "/gifts")
}
