package route

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.Render("posts", fiber.Map{})
	})

	app.Get("/post/:slug", func(c *fiber.Ctx) error {
		return c.Render("post-detail", fiber.Map{})
	})
}
