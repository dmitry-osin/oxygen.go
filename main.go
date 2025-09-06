package main

import (
	"log"
	"oxygenBlog/config"
	"oxygenBlog/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		return
	}

	engine := handlebars.New("./public/views", ".hbs")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views:         engine,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       configuration.ApplicationName,
	})
	app.Static("/static", "./public/static")

	route.RegisterPublicRoutes(app)
	route.RegisterAdminRoutes(app)

	err = app.Listen(configuration.Port)
	if err != nil {
		log.Fatal(err)
	}
}
