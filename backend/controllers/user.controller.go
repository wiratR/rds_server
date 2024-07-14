package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbconn "github.com/wiratR/go-orm-jwt/database"
	"github.com/wiratR/go-orm-jwt/models"
	"github.com/wiratR/go-orm-jwt/utils"
)

// get all todos
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	dbconn.DB.Find(&users)
	return c.Status(200).JSON(users)
}

// GetMe func get current user
// @Description Get current user.
// @Summary get current user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users/me [get]
func GetMe(c *fiber.Ctx) error {
	id := c.Cookies("user_id")
	var user models.User
	result := dbconn.DB.Find(&user, uuid.MustParse(id))

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User id " + id + " not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": gin.H{"user": models.FilterUserRecord(&user)}})
}

// Get User by Id func get user infomation by id
// @Description get user infomation by id
// @Summary get user infomation by id
// @Tags Uer
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.User
// @Router /users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	var isSuccess bool = false
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	var user models.User

	result := dbconn.DB.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
	}
	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"user": gin.H{"user": user}}))
}

// Update user detail
// @Description update user detail
// @Summary update user detail
// @Tags Uer
// @Accept json
// @Produce json
// @Param Payload body models.UserUpdate true "User Update Data"
// @Success 200 {array} models.User
// @Router /users/update [patch]
func UpdateUser(c *fiber.Ctx) error {
	var payload *models.UserUpdate
	var isSuccess bool = false

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
	}

	var oldUser models.User
	id := c.Cookies("user_id")

	result := dbconn.DB.First(&oldUser, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the user ID is not found"}))
	}

	var hashedPassword string
	// not update password
	if payload.Password == "" {
		hashedPassword = oldUser.Password
	} else {
		hashedPassword = utils.HashPassword(payload.Password)
	}

	updateUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
		Role:     &payload.Role,
		Verified: &payload.Verified,
	}

	result = dbconn.DB.Model(&models.User{}).Where("id = ?", id).Updates(&updateUser)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid email or password"}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"data": gin.H{"user": models.FilterUserRecord(&updateUser)}}))
}

func DeleteUser(c *fiber.Ctx) error {
	var isSuccess bool = false
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}
	// Delete a record by a condition
	result := dbconn.DB.Where("id = ?", id).Delete(&models.User{})

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"status": "fail", "message": "Delete failed"}))
	}
	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Delete user is success"})
}
