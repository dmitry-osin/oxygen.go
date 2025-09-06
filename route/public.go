package route

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{})
	})

	app.Post("/contact", func(c *fiber.Ctx) error {
		// Handle contact form submission
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Thank you for your message! We'll get back to you soon.",
		})
	})

	app.Get("/post/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		return c.Render("post", fiber.Map{
			"slug": slug,
		})
	})
}
