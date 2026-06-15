package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/payment-backend/config"
	routes "github.com/payment-backend/route"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

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

// WAJIB huruf besar
func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
