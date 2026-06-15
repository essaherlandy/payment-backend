package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/payment-backend/config"
	routes "github.com/payment-backend/route"
)

var app *fiber.App

func init() {

	app = fiber.New()

	fmt.Println("FIBER CREATED")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	config.ConfigDatabase()

	fmt.Println("DB CONNECTED")

	routes.Setup(app)

	fmt.Println("ROUTES READY")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
