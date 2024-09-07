package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	"github.com/wiratR/rds_server/utils"
	"gorm.io/gorm"
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
// @Success 200 {array} models.ResponseSuccessUser
// @Router /users/me [get]
func GetMe(c *fiber.Ctx) error {
	id := c.Cookies("user_id")

	fmt.Printf("user controller : id = %s\n", id)

	var user models.User
	result := dbconn.DB.Find(&user, uuid.MustParse(id))

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User id " + id + " not found"})
	}

	accountIDPtr, err := GetAccountIDByUserId(*user.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User id " + id + " not found"})
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": gin.H{"user": models.FilterUserRecord(&user, accountIDPtr)}})
}

// Get User by Id func get user infomation by id
// @Description get user infomation by id
// @Summary get user infomation by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.ResponseSuccessUser
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

	accountIDPtr, err := GetAccountIDByUserId(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"user": models.FilterUserRecord(&user, accountIDPtr)}))
}

// Update user detail
// @Description update user detail
// @Summary update user detail
// @Tags User
// @Accept json
// @Produce json
// @Param Payload body models.UserUpdate true "User Update Data"
// @Success 200 {array} models.ResponseSuccessUser
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
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     strings.ToLower(payload.Email),
		Phone:     payload.Phone,
		Password:  hashedPassword,
		Role:      &payload.Role,
		Verified:  &payload.Verified,
	}

	result = dbconn.DB.Model(&models.User{}).Where("id = ?", id).Updates(&updateUser)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid email or password"}))
	}

	accountIDPtr, err := GetAccountIDByUserId(*oldUser.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess,
		fiber.Map{"data": gin.H{"user": models.FilterUserRecord(&updateUser, accountIDPtr)}}))
}

// Update user detail by user id
// @Description update user detail by user id
// @Summary update user detail by user id
// @Tags User
// @Accept json
// @Produce json
// @Param Payload body models.UserUpdate true "User Update Data"
// @Success 200 {array} models.ResponseSuccessUser
// @Router /users/update/{id} [patch]
func UpdateUserById(c *fiber.Ctx) error {
	var payload *models.UserUpdate
	var isSuccess bool = false

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
	}

	var oldUser models.User

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
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     strings.ToLower(payload.Email),
		Phone:     payload.Phone,
		Password:  hashedPassword,
		Role:      &payload.Role,
		Verified:  &payload.Verified,
	}

	result = dbconn.DB.Model(&models.User{}).Where("id = ?", id).Updates(&updateUser)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid email or password"}))
	}

	accountIDPtr, err := GetAccountIDByUserId(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"data": gin.H{"user": models.FilterUserRecord(&updateUser, accountIDPtr)}}))
}

// Delete user by user id
// @Description delete user by user id
// @Summary Delete user by user id
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.ResponseError "success"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /users/delete/{id} [delete]
func DeleteUserById(c *fiber.Ctx) error {
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

// Update user password
// @Description update user password
// @Summary update user password
// @Tags User
// @Accept json
// @Produce json
// @Param Payload body models.UserPasswordUpdate true "Update user password"
// @Success 200 {array} models.ResponseSuccessUser
// @Router /users/updatepassword/{id} [patch]
func UpdateUserPasswordById(c *fiber.Ctx) error {
	var isSuccess bool = false
	var payload *models.UserPasswordUpdate
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
	}

	var oldUser models.User
	result := dbconn.DB.First(&oldUser, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the user ID is not found"}))
	}

	if err := utils.ComparePassword(oldUser.Password, payload.OldPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid Password"}))
	}

	var hashedPassword = utils.HashPassword(payload.NewPassword)

	updateUser := models.User{
		UserName:  oldUser.UserName,
		FirstName: oldUser.FirstName,
		LastName:  oldUser.LastName,
		Email:     strings.ToLower(oldUser.Email),
		Phone:     oldUser.Phone,
		Password:  hashedPassword,
		Role:      oldUser.Role,
		Verified:  oldUser.Verified,
	}

	result = dbconn.DB.Model(&models.User{}).Where("id = ?", id).Updates(&updateUser)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Invalid email or password"}))
	}

	accountIDPtr, err := GetAccountIDByUserId(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("account not found")
		} else {
			fmt.Println("Error occurred:", err)
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err}))
		}
	} else {
		fmt.Printf("account ID: %v\n", *accountIDPtr)
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"data": gin.H{"user": models.FilterUserRecord(&updateUser, accountIDPtr)}}))
}
