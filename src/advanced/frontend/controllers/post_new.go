package controllers

import (
	"html/template"
	"path"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/gem"
	"net/http"
	"strings"
	"time"
)

type PostNew struct {
	Controller
	tmpl *template.Template
}

func (c *PostNew) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("post", "new"))

	return nil
}

func (c *PostNew) GET(ctx *gem.Context) {
	if c.IsGuest(ctx) {
		ctx.Redirect("/login", http.StatusFound)
		return
	}

	c.RenderTemplate(c.tmpl, ctx, nil)
}

func (c *PostNew) POST(ctx *gem.Context) {
	user := c.GetUser(ctx)
	if user == nil {
		c.JSON(ctx, 10, "please sign in and then try again", nil)
		return
	}

	title := strings.TrimSpace(ctx.PostFormValue("title"))
	description := strings.TrimSpace(ctx.PostFormValue("description"))
	content := strings.TrimSpace(ctx.PostFormValue("content"))
	if title == "" {
		c.JSON(ctx, 100020, "title should not be blank", nil)
		return
	}
	if description == "" {
		c.JSON(ctx, 100021, "description should not be blank", nil)
		return
	}
	if content == "" {
		c.JSON(ctx, 100022, "content should not be blank", nil)
		return
	}

	now := time.Now()

	post := models.Post{
		UserID:      user.ID,
		Title:       title,
		Description: description,
		Content:     content,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := c.db.Save(&post).Error; err != nil {
		ctx.Logger().Errorf("failed to create a new post: %s\n", err)
		c.JSON(ctx, 100022, "failed to create a new post", nil)
		return
	}

	c.JSON(ctx, 0, "post successfully", post.ID)
}
