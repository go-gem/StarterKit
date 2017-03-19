package controllers

import (
	"html/template"
	"path"

	"github.com/go-gem/gem"
)

type Index struct {
	Controller
	tmpl *template.Template
}

func (c *Index) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("site", "index"))

	return nil
}

func (c *Index) GET(ctx *gem.Context) {
	c.RenderTemplate(c.tmpl, ctx, nil)
}
