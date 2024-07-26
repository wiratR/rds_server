package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
	"github.com/wiratR/rds_server/middleware"
)

func UserRoute(route fiber.Router) {
	// routes "/api/users"
	route.Get("/me", middleware.DeserializeUser, controllers.GetMe)
	route.Get("/:id", middleware.DeserializeUser, controllers.GetUserById)
	route.Patch("/update", middleware.DeserializeUser, controllers.UpdateUser)
	route.Delete("/deactive", middleware.DeserializeUser, controllers.DeleteUser)
}
