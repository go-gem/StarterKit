package controllers

import (
	"github.com/go-gem/gem"
	"net/http"
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
