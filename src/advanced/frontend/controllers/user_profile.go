package controllers

import (
	"html/template"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/gem"
	"path"
)

type UserProfile struct {
	Controller

	tmpl *template.Template
}

func (c *UserProfile) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("user", "profile"))

	return nil
}

func (c *UserProfile) GET(ctx *gem.Context) {
	var profile models.User

	name, err := gem.String(ctx.UserValue("name"))
	if err == nil && name != "" {
		c.db.Where("username = ?", name).Find(&profile)
	}

	c.RenderTemplate(c.tmpl, ctx, map[string]interface{}{
		"Profile": profile,
	})
}
