package mains

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/payment-backend/config"
	routes "github.com/payment-backend/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	godotenv.Load(".env")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	config.ConfigDatabase()

	routes.Setup(app)

	app.Listen(":8000")
}
