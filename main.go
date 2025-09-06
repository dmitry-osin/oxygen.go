package main

import (
	"log"
	"oxygenBlog/config"
	"oxygenBlog/domain"
	"oxygenBlog/middleware"
	"oxygenBlog/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connection, err := domain.NewConnection(*configuration)
	if err != nil {
		log.Fatal(err)
	}

	connection.RegisterModel(
		&domain.User{},
		&domain.PostsTags{},
		&domain.Tag{},
		&domain.Post{},
	)

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

	app.Use(middleware.MinifyCSS())
	app.Static("/static", "./public/static")

	route.RegisterPublicRoutes(app)
	route.RegisterAdminRoutes(app)

	err = app.Listen(configuration.Port)
	if err != nil {
		log.Fatal(err)
	}
}
