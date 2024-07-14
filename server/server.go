package main

import (
	"log"
	"os"

	"github.com/MeongGanas/golang-e-commerce/database"
	"github.com/MeongGanas/golang-e-commerce/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	routes.AllRoutes(app)

	log.Fatal(app.Listen(":" + PORT))
}
