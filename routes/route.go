package routes

import (
	"apikamar/controllers"
	"apikamar/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	kamar := r.Group("/api")

	kamar.Get("/", middlewares.AuthMiddleware, controllers.Index)
	kamar.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	kamar.Post("/", middlewares.AuthMiddleware, controllers.Create)
	kamar.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
	kamar.Delete("/:id", middlewares.AuthMiddleware, controllers.Delete)
}