package controllers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/wiratR/rds_server/config"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	"github.com/wiratR/rds_server/utils"
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
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     strings.ToLower(payload.Email),
		Phone:     payload.Phone,
		Password:  hashedPassword,
	}

	// create new user
	result := dbconn.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Passwords do not match"}))
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Something bad happened"}))
	}

	accountIDPtr, err := GetAccountIDByUserId(*newUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "User ID alredy use in accout"}))
	}

	isSuccess = true
	// return success 201
	return c.Status(fiber.StatusCreated).JSON(models.ApiResponse(isSuccess, fiber.Map{"user": models.FilterUserRecord(&newUser, accountIDPtr)}))
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

// SignInByPhone func Login by phone number
// @Description Login by phone number
// @Summary Login by phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param Payload body models.SignInByPhone true "Login Data"
// @Success 200 {object} models.SignInResponse "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /auth/loginbyphone [post]
func SignInByPhone(c *fiber.Ctx) error {
	log.Println("auth controller : start signin ")

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var payload *models.SignInByPhone
	var isSuccess bool = false

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	fmt.Printf("auth controller : phone = %s\n", payload.Phone)
	fmt.Printf("auth controller : password = %s\n", payload.Password)

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
		//return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": errors})
	}

	var user models.User
	result := dbconn.DB.First(&user, "phone = ?", strings.ToLower(payload.Phone))
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

	id := c.Cookies("user_id")

	fmt.Printf("auth controller check: id = %s\n", id)

	accountIDPtr, err := GetAccountIDByUserId(*user.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	}
	// } else {
	// 	fmt.Printf("account ID: %v\n", *accountIDPtr)
	// 	return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "User ID alredy use in accout"}))
	// }

	isSuccess = true
	// return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"token": access_token}))
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, models.SignInResponse{
		Token:       access_token,
		UserDetails: models.FilterUserRecord(&user, accountIDPtr),
	}))

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
