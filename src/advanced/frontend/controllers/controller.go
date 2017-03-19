package controllers

import (
	"fmt"
	"html/template"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/gem"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

// Controller base controller
type Controller struct {
	gem.WebController
	app           *gem.Application
	sessionsStore sessions.Store
	db            *gorm.DB
}

func (c *Controller) SessionStore() sessions.Store {
	return c.sessionsStore
}

func (c *Controller) Init(app *gem.Application) error {
	c.app = app

	var ok bool

	if c.sessionsStore, ok = app.Component("sessionsStore").(sessions.Store); !ok {
		return fmt.Errorf("invalid sessions store component")
	}

	if c.db, ok = app.Component("db").(*gorm.DB); !ok {
		return fmt.Errorf("invalid database component")
	}

	return nil
}

func (c *Controller) Render(filenames ...string) *template.Template {
	return template.Must(c.app.Templates().Render("main", filenames...))
}

func (c *Controller) RenderTemplate(tmpl *template.Template, ctx *gem.Context, data map[string]interface{}) {
	v := map[string]interface{}{
		"User": c.GetUser(ctx),
	}

	for key, value := range data {
		v[key] = value
	}

	tmpl.Execute(ctx, v)
}

func (c *Controller) JSON(ctx *gem.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func (c *Controller) GetSession(ctx *gem.Context) (session *sessions.Session) {
	session, _ = c.sessionsStore.Get(ctx.Request, "GOSESSION")
	return
}

func (c *Controller) SaveSession(session *sessions.Session, ctx *gem.Context) {
	session.Save(ctx.Request, ctx.Response)
}

func (c *Controller) GetUser(ctx *gem.Context) *models.User {
	session := c.GetSession(ctx)
	if v, ok := session.Values["user"].(models.User); ok {
		return &v
	}

	return nil
}

func (c *Controller) IsGuest(ctx *gem.Context) bool {
	user := c.GetUser(ctx)
	if user != nil {
		return false
	}

	return true
}
