package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	"github.com/wiratR/rds_server/utils"
)

// CreateAccount func create new account with user id
// @Description create new account with user id
// @Summary create new account with user id
// @Tags Accounts
// @Accept json
// @Produce json
// @Param Payload body models.AccoutCreateInput true "Account Data"
// @Success 201 {array} models.ResponseSuccessAccount "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /accounts/create [post]
func CreateAccount(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var payload *models.AccoutCreateInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}
	// validation message payload
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": errors}))
	}
	// validation paload userId
	userId, err := uuid.Parse(payload.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}
	var user models.User
	findUserIDResult := dbconn.DB.Find(&user, userId)
	if findUserIDResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
	}
	// genarte new access token
	accountToken := utils.GenerateAccountToken()
	newAccount := models.Account{
		AccountToken:   accountToken,
		AccountType:    "mobile",
		Status:         "bining",
		Balance:        0,
		BlockFlag:      0,
		LastEntrySpId:  0,
		LastEntryLocId: 0,
		Active:         true,
		UserID:         &userId,
		User:           &user,
	}
	// create new user
	createAccoutResult := dbconn.DB.Create(&newAccount)
	if createAccoutResult.Error != nil && strings.Contains(createAccoutResult.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "duplicate key value violates unique"}))
	} else if createAccoutResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Something bad happened"}))
	}
	isSuccess = true
	// return success 201
	return c.Status(fiber.StatusCreated).JSON(models.ApiResponse(isSuccess, fiber.Map{"account": models.FilterAccountRecord(&newAccount)}))
}

func GetAccountIDByUserId(userId uuid.UUID) (*uuid.UUID, error) {

	if dbconn.DB == nil {
		return nil, errors.New("database connection is nil")
	}

	var account models.Account
	result := dbconn.DB.Where("user_id = ?", userId).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return (*uuid.UUID)(account.ID), nil
}

// Get Account by user id func get account infomation by userid
// @Description get account by user id func get account infomation by userid
// @Summary get account by user id func get account infomation by userid
// @Tags Accounts
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {array} models.ResponseSuccessAccount
// @Router /accounts/{userId} [get]
func GetAccountByUserId(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	userId, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	var user models.User
	userResult := dbconn.DB.Find(&user, userId)

	if userResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User id not found"})
	}

	var account models.Account
	result := dbconn.DB.Where("user_id = ?", userId).First(&account)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
	}

	account.User = &user

	account.User.ID = user.ID
	account.User.FirstName = user.FirstName

	fmt.Println(account.User.ID)

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"account": models.FilterAccountRecord(&account)}))

}

// func UpdateAccount(c *fiber.Ctx) error {
// 	if dbconn.DB == nil {
// 		return errors.New("database connection is nil")
// 	}
// 	var isSuccess bool = false
// }

// func DeactiveAccount(c *fiber.Ctx) error {
// 	if dbconn.DB == nil {
// 		return errors.New("database connection is nil")
// 	}
// 	var isSuccess bool = false
// }
