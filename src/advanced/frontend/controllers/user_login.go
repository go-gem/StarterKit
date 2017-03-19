package controllers

import (
	"fmt"
	"html/template"
	"path"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/StarterKit/src/advanced/common/validators"
	"github.com/go-gem/gem"
	"net/http"
	"strings"
)

type UserLogin struct {
	Controller
	tmpl *template.Template
}

func (c *UserLogin) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("user", "login"))

	return nil
}

func (c *UserLogin) GET(ctx *gem.Context) {
	if !c.IsGuest(ctx) {
		ctx.Redirect("/", http.StatusFound)
		return
	}

	c.RenderTemplate(c.tmpl, ctx, nil)
}

func (c *UserLogin) POST(ctx *gem.Context) {
	account := strings.TrimSpace(ctx.PostFormValue("account"))
	password := strings.TrimSpace(ctx.PostFormValue("password"))

	if account == "" {
		c.JSON(ctx, 100020, "account should not be blank", nil)
		return
	}
	if password == "" {
		c.JSON(ctx, 100020, "password should not be blank", nil)
		return
	}

	var user models.User
	if validators.IsEmail(account) {
		c.db.Where("email = ?", account).Find(&user)
	} else {
		c.db.Where("username = ?", account).Find(&user)
	}

	if c.db.NewRecord(user) {
		c.JSON(ctx, 100020, fmt.Sprintf("the account named %q does not exist", account), nil)
		return
	}

	if !user.ValidatePassword(password) {
		c.JSON(ctx, 100020, "invalid account or password", nil)
		return
	}

	session := c.GetSession(ctx)
	// Store user info into session.
	session.Values["user"] = user
	if err := session.Save(ctx.Request, ctx.Response); err != nil {
		ctx.Logger().Errorf("login failure: %s\n", err)
		c.JSON(ctx, 0, "login failure", nil)
		return
	}

	c.JSON(ctx, 0, "login successfully", nil)
}
