package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payment-backend/controllers"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/register",
		controllers.Register,
	)

	api.Post("/login",
		controllers.Login,
	)

}
