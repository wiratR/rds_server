package controllers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/go-orm-jwt/config"
	dbconn "github.com/wiratR/go-orm-jwt/database"
	"github.com/wiratR/go-orm-jwt/models"
	"github.com/wiratR/go-orm-jwt/utils"
)

// SignUpUser func Register New Account
// @Description register new account
// @Summary register new account
// @Tags Auth
// @Accept json
// @Produce json
// @Param Payload body models.SignUpInput true "Register Data"
// @Success 201 {array} models.ResponseSuccessUser "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /auth/register [post]
func SignUpUser(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var payload *models.SignUpInput
	var isSuccess bool = false

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}
	// validation payload
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
	}
	// validation Password and PasswordConfirm
	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Passwords do not match"}))
	}
	// create Hashed Password
	hashedPassword := utils.HashPassword(payload.Password)

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
	}

	// create new user
	result := dbconn.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Passwords do not match"}))
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Something bad happened"}))
	}

	isSuccess = true
	// return success 201
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"user": models.FilterUserRecord(&newUser)}))
}

// SignInUser func Login
// @Description login
// @Summary login
// @Tags Auth
// @Accept json
// @Produce json
// @Param Payload body models.SignInInput true "Login Data"
// @Success 200 {object} models.ResponseSuccessToken "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /auth/login [post]
func SignInUser(c *fiber.Ctx) error {

	log.Println("auth controller : start signin ")

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var payload *models.SignInInput
	var isSuccess bool = false

	// fmt.Printf("auth controller : email = %s\n", payload.Email)
	// fmt.Printf("auth controller : password = %s\n", payload.Password)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	fmt.Printf("auth controller : email = %s\n", payload.Email)
	fmt.Printf("auth controller : password = %s\n", payload.Password)

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
		//return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": errors})
	}

	var user models.User
	result := dbconn.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid email or Password"}))
	}

	if err := utils.ComparePassword(user.Password, payload.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid Password"}))
	}

	config, _ := config.LoadConfig(".")

	// Generate Token
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Generate Refresh Token Failed"}))
	}
	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    access_token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.DomainHost,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshtoken",
		Value:    refresh_token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.DomainHost,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.DomainHost,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    user.ID.String(),
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   config.DomainHost,
	})

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"token": access_token}))
}

// LogoutUser func Logout
// @Description logout
// @Summary logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseError "Ok"
// @Router /auth/logout [get]
func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(true, fiber.Map{"message": "Logout success"}))
}
