package domain

import (
	"github.com/citrusinesis/thread/pkg/domain/entity"
	"github.com/gofiber/fiber/v2"
	"sync"
)

var routerOnce sync.Once

type Group interface {
	SetRoute(*fiber.App)
}

type Router struct {
	groups []Group
}

func NewRouter(entity *entity.Group) (router *Router) {
	// Should add parameter of each group
	routerOnce.Do(func() {
		router = &Router{
			[]Group{
				// TODO: Add router groups.
				entity,
			},
		}
	})
	return
}

func (r *Router) RegisterRoute(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API By Thread🪡, Powered by Fiber🚀; Wired🛜")
	})

	for _, group := range r.groups {
		group.SetRoute(app)
	}
}
