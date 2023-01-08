package main

import (
	core "github.com/DisposableChat/api-core"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	router fiber.Router
}

func (r *Router) Configure() {
	r.SetRoutes()
}

func (r *Router) SetRoutes()  {
	r.SetIndividual()
}

func (r *Router) SetIndividual() {
	r.router.Get("/", func(c *fiber.Ctx) error {
		message, data, err := GetSession(c)
		if err != nil {
			return core.Error(c, err.Error(), nil)
		}

		return core.OK(c, message, data)
	})

	r.router.Post("/", func(c *fiber.Ctx) error {
		message, data, err := CreateSession(c)
		if err != nil {
			return core.Error(c, err.Error(), nil)
		}

		return core.OK(c, message, data)
	})
}
