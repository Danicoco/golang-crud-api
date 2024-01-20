package main

import (
	"crud-with-postgres/configs"
	"crud-with-postgres/database"
	"crud-with-postgres/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run(app *fiber.App) error {
	app.Use(cors.New())
	configs.LoadEnvs()
	routes.Routes(app)

	_, err := database.DBConn()

	if err != nil {
		panic(err)
	}

	return nil
}
