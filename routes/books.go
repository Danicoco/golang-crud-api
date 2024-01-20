package routes

import (
	"crud-with-postgres/controllers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(v1 fiber.Router) {
	v1.Get("/Books", controllers.FetchBook)
	v1.Get("/Books/:id", controllers.GetBook)
	v1.Post("/Books", controllers.CreateBook)
	v1.Put("/Books/:id", controllers.UpdateBook)
	v1.Delete("/Books/:id", controllers.DeleteBook)
}
