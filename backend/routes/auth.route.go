package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/go-orm-jwt/controllers"
	"github.com/wiratR/go-orm-jwt/middleware"
)

func AuthRoute(route fiber.Router) {
	// routes "/api/auth"
	route.Post("/register", controllers.SignUpUser)
	route.Post("/login", controllers.SignInUser)
	route.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
}
