package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
)

func PaymentRoute(route fiber.Router) {
	// routes "/api/payment"
	route.Post("/nativepay", controllers.NativePay)
}
