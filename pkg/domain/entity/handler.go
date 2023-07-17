package entity

import (
	"sync"
)

var handlerOnce sync.Once

type Handler struct {
	service *Service
}

func NewItemHandler(service *Service) (handler *Handler) {
	handlerOnce.Do(func() {
		handler = &Handler{service}
	})
	return
}

/*
	TODO: Define handlers.

	Example:
		func (h *Handler) HANDLER(c *fiber.Ctx) error {
			value, err := h.service.SERVICE()
			if err != nil {
				return err
			}

			return c.JSON(value)
		}
*/
