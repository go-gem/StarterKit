package controllers

import (
	"html/template"
	"path"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/gem"
)

type PostDetail struct {
	Controller
	tmpl *template.Template
}

func (c *PostDetail) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("post", "detail"))

	return nil
}

func (c *PostDetail) GET(ctx *gem.Context) {
	var id int
	if num, err := gem.Int(ctx.UserValue("id")); err == nil && num > 0 {
		id = num
	}

	post := models.Post{}
	c.db.Where("id = ?", id).Find(&post)
	c.db.Model(&post).Related(&post.User)

	c.RenderTemplate(c.tmpl, ctx, map[string]interface{}{
		"Post": post,
	})
}
