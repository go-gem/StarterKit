package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-gem/StarterKit/src/advanced/common/models"
	"github.com/go-gem/StarterKit/src/advanced/common/validators"
	"github.com/go-gem/gem"
)

type UserJoin struct {
	Controller

	tmpl *template.Template
}

func (c *UserJoin) Init(app *gem.Application) error {
	c.tmpl = c.Render(path.Join("user", "join"))

	return nil
}

func (c *UserJoin) GET(ctx *gem.Context) {
	if !c.IsGuest(ctx) {
		ctx.Redirect("/", http.StatusFound)
		return
	}

	c.RenderTemplate(c.tmpl, ctx, nil)
}

func (c *UserJoin) POST(ctx *gem.Context) {
	email := strings.TrimSpace(ctx.PostFormValue("email"))
	username := strings.TrimSpace(ctx.PostFormValue("username"))
	password := strings.TrimSpace(ctx.PostFormValue("password"))

	var err error

	// validate
	if err = c.checkEmail(email); err != nil {
		c.JSON(ctx, 100010, err.Error(), nil)
		return
	}
	if err = c.checkUsername(username); err != nil {
		c.JSON(ctx, 100011, err.Error(), nil)
		return
	}
	if err = c.checkPassword(password); err != nil {
		c.JSON(ctx, 100012, err.Error(), nil)
		return
	}

	// save
	now := time.Now()
	user := &models.User{
		Email:     email,
		Username:  username,
		CreatedAt: now,
		UpdatedAt: now,
	}

	user.GeneratePassword(password)

	if err = c.db.Save(&user).Error; err != nil {
		ctx.Logger().Errorf("failed to create user: %s", err)
		c.JSON(ctx, 100013, "failed to create user", nil)
		return
	}

	c.JSON(ctx, 0, "join successfully", nil)
}

var errEmptyEmail = errors.New("email should not be blank")

func (c *UserJoin) checkEmail(email string) error {
	if email == "" {
		return errEmptyEmail
	}

	if !validators.IsEmail(email) {
		return fmt.Errorf("invalid email: %s", email)
	}

	var user models.User
	c.db.Where("email = ?", email).Find(&user)

	if !c.db.NewRecord(user) {
		return fmt.Errorf("the email has been used: %s", email)
	}

	return nil
}

var errEmptyUsername = errors.New("username should not be blank")

func (c *UserJoin) checkUsername(username string) error {
	if username == "" {
		return errEmptyUsername
	}

	var user models.User
	c.db.Where("username = ?", username).Find(&user)

	if !c.db.NewRecord(user) {
		return fmt.Errorf("the username has been used: %s", username)
	}

	return nil
}

var errInvalidPassword = errors.New("invalid password")

func (c *UserJoin) checkPassword(password string) error {
	if password == "" || len(password) < 8 {
		return errInvalidPassword
	}

	return nil
}
