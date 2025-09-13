package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/transport/http/handler"
)

func Register(app *fiber.App, ch *handler.CustomerHandler) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error { return c.SendString("ok") })

	g := api.Group("/customers")
	g.Get("/", ch.List)
	g.Get("/:id", ch.Get)
	g.Post("/", ch.Create)
	g.Put("/:id", ch.Update)
	g.Delete("/:id", ch.Delete)
}
