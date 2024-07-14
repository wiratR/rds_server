package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/go-orm-jwt/config"
	"github.com/wiratR/go-orm-jwt/database"
	"github.com/wiratR/go-orm-jwt/models"
	"github.com/wiratR/go-orm-jwt/utils"
)

func DeserializeUser(c *fiber.Ctx) error {

	var access_token string

	authorizationHeader := c.Get("Authorization")

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		access_token = strings.TrimPrefix(authorizationHeader, "Bearer ")
	} else if c.Cookies("token") != "" {
		access_token = c.Cookies("token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiResponse(false, fiber.Map{"message": "You are not logged in"}))
	}

	config, _ := config.LoadConfig(".")
	sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(false, fiber.Map{"message": fmt.Sprintf("invalidate token: %v", err)}))
	}

	var user models.User
	database.DB.First(&user, "id = ?", fmt.Sprint(sub))

	if user.ID.String() != sub {
		return c.Status(fiber.StatusForbidden).JSON(models.ApiResponse(false, fiber.Map{"message": "the user belonging to this token no logger exists"}))
	}

	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
