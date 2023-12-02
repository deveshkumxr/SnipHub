package main

import (
	"os"

	"github.com/deveshkumxr/SnipHub/db"
	"github.com/deveshkumxr/SnipHub/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	app := fiber.New()
	db.Connect()
	routes.InitRoutes(app)

	app.Listen(port)

}