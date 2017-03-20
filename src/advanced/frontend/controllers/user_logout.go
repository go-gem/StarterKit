package controllers

import (
	"net/http"

	"github.com/go-gem/gem"
)

type UserLogout struct {
	Controller
}

func (c *UserLogout) Methods() []string {
	return []string{gem.MethodPost}
}

func (c *UserLogout) POST(ctx *gem.Context) {
	if !c.IsGuest(ctx) {
		session := c.GetSession(ctx)
		delete(session.Values, "user")
		session.Save(ctx.Request, ctx.Response)
	}

	ctx.Redirect("/login", http.StatusFound)
}
