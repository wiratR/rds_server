package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
	"github.com/wiratR/rds_server/middleware"
)

func TxnHistoryRoute(route fiber.Router) {
	// routes "/api/txnhistories"
	// route.Post("/create", middleware.DeserializeUser, controllers.CreateTxnHistory)
	route.Get("/:accountId", middleware.DeserializeUser, controllers.GetTxnHistoryByAccountId)
}
