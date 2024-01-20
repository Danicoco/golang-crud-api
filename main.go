package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	err := Run(app)

	if err != nil {
		panic(err)
	}

	app.Listen(":3000")
}
