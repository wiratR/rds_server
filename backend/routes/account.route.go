package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
	"github.com/wiratR/rds_server/middleware"
)

func AccountRoute(route fiber.Router) {
	// routes "/api/accounts"
	route.Post("/create", middleware.DeserializeUser, controllers.CreateAccount)
	route.Get("/:userId", middleware.DeserializeUser, controllers.GetAccountByUserId)
}
