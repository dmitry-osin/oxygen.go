package route

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}

func About(ctx *fiber.Ctx) error {
	return ctx.Render("about", fiber.Map{})
}

func Contact(ctx *fiber.Ctx) error {
	return ctx.Render("contact", fiber.Map{})
}

func Blog(ctx *fiber.Ctx) error {
	return ctx.Render("blog", fiber.Map{})
}
