package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
	"github.com/wiratR/rds_server/middleware"
)

func AuthRoute(route fiber.Router) {
	// routes "/api/auth"
	route.Post("/register", controllers.SignUpUser)
	route.Post("/login", controllers.SignInUser)
	route.Post("/loginbyphone", controllers.SignInByPhone)
	route.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
}
