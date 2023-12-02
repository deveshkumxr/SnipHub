package routes

import (
	"github.com/deveshkumxr/SnipHub/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	snippets := app.Group("snippets")

	snippets.Get("/", controllers.GetSnippets)
	snippets.Post("/", controllers.CreateSnippet)
	snippets.Put("/:snippet_id", controllers.UpdateSnippet)
	snippets.Delete("/:snippet_id", controllers.DeleteSnippet)
}