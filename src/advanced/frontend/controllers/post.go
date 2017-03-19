package controllers

import (
	"html/template"
	"path"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/gem"
)

type Post struct {
	Controller
	tmpl *template.Template
}

func (c *Post) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("post", "index"))

	return nil
}

func (c *Post) GET(ctx *gem.Context) {
	pageSize := 10
	pageNum := 1

	if num, err := gem.Int(ctx.UserValue("page")); err == nil && num > 0 {
		pageNum = num
	}

	posts := make([]models.Post, 10)
	c.db.Order("id DESC").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&posts)

	c.RenderTemplate(c.tmpl, ctx, map[string]interface{}{
		"Posts": posts,
	})
}
