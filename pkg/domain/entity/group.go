package entity

import (
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Group struct {
	handler *Handler
}

var routerOnce sync.Once

func NewItemGroup(handler *Handler) (group *Group) {
	routerOnce.Do(func() {
		group = &Group{handler}
	})
	return
}

func (r *Group) SetRoute(app *fiber.App) {
	/*
		TODO: Define routes with form of group.

		Example:
		app.Group("/entity").
			Get("/", r.handler.HANDLER)
	*/
	panic("UNDEFINED")
}
