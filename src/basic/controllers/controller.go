package controllers

import (
	"html/template"

	"github.com/go-gem/gem"
)

// Controller base controller
type Controller struct {
	gem.WebController
	app *gem.Application
}

// Init initialize  controller.
func (c *Controller) Init(app *gem.Application) error {
	c.app = app

	return nil
}

// Render template via main layout.
func (c *Controller) Render(filenames ...string) (*template.Template, error) {
	return c.app.Templates().Render("main", filenames...)
}
