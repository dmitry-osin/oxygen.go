package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	"oxygen.go/route"
	"oxygen.go/route/area"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Oxygen Blogging Application",
		Views:   pug.New("./views", ".pug"),
	})

	app.Static("/", "./public")

	app.Get("/", route.Index)
	app.Get("/about", route.About)
	app.Get("/contact", route.Contact)

	app.Get("/admin", area.Admin)
	app.Get("/admin/posts", area.Posts)
	app.Get("/admin/post/:url", area.Post)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
