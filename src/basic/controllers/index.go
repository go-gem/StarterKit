package controllers

import (
	"html/template"

	"github.com/go-gem/gem"
)

// Index controller
type Index struct {
	Controller
	tmpl *template.Template
}

// Init initialize controller
func (c *Index) Init(app *gem.Application) (err error) {
	c.tmpl, err = c.Render("index")
	if err != nil {
		return err
	}

	return nil
}

func (c *Index) GET(ctx *gem.Context) {
	c.tmpl.Execute(ctx, nil)
}
