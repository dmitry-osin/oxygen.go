package area

import "github.com/gofiber/fiber/v2"

func Admin(ctx *fiber.Ctx) error {
	return ctx.Render("area/admin", fiber.Map{})
}

func Posts(ctx *fiber.Ctx) error {
	return ctx.Render("area/posts", fiber.Map{})
}

func Post(ctx *fiber.Ctx) error {
	return ctx.Render("area/post", fiber.Map{})
}
