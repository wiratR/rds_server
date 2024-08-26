package controllers

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
)

// Convert TxnHistory to TxnHistoryResponse
func toTxnHistoryResponse(txn models.TxnHistory) models.TxnHistoryResponse {

	LocEntryName := getStationNameById(txn.LocEntryId)
	LocExitName := getStationNameById(txn.LocExitId)
	LineIdOfEntry := getLineIdByStationId(txn.LocEntryId)
	LineIdOfExit := getLineIdByStationId(txn.LocEntryId)
	LineNameOfEntry := getLineNameByLineId(LineIdOfEntry)
	LineNameOfExit := getLineNameByLineId(LineIdOfExit)

	return models.TxnHistoryResponse{
		TxnRefId:  txn.TxnRefId,
		TxnAmount: txn.TxnAmount,
		TxnDetail: models.TxnDetail{TxnTypeId: txn.TxnTypeId}, // Populate TxnTypeName if necessary
		SpDetail:  models.SpDetail{SpId: txn.SpId},            // Populate SpName if necessary
		LocEntryDetail: models.LocDetail{
			LocId:   txn.LocEntryId,
			LocName: LocEntryName,
			LineDetail: models.LineDetail{
				LineId:   LineIdOfEntry,
				LineName: LineNameOfEntry,
			},
		},
		LocExitDetail: models.LocDetail{
			LocId:   txn.LocExitId,
			LocName: LocExitName,
			LineDetail: models.LineDetail{
				LineId:   LineIdOfExit,
				LineName: LineNameOfExit,
			},
		},
		EquipmentNumber: txn.EquipmentNumber,
		CreatedAt:       txn.CreatedAt,
		UpdatedAt:       txn.UpdatedAt,
	}
}

// GetTxnHistory retrieves a transaction history by account ID
// @Description get account by user id func get account infomation by userid
// @Summary get account by user id func get account infomation by userid
// @Tags Txnhistories
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {array} models.TxnHistoryResponse
// @Router /accounts/{accountId} [get]
func GetTxnHistoryByAccountId(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	accountId, err := uuid.Parse(c.Params("accountId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	var txn models.TxnHistory

	if err := dbconn.DB.Preload("Account").First(&txn, "account_id = ?", accountId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the account ID then transaction not found"}))
	}

	response := toTxnHistoryResponse(txn)

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"txn_history": response}))
	// return c.JSON(response)
}

// CreateTxnHistory
// @Description create transaction history with account id
// @Summary create transaction history with account id
// @Tags TxnHistory
// @Accept json
// @Produce json
// @Param Payload body models.TxnCreateInput true "Txn Create Input Data"
// @Success 201 {array} models.ResponseSuccess "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /txnhistories/create [post]
func CreateTxnHistory(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var payload *models.TxnCreateInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	newTxn := models.TxnHistory{
		ID:        uuid.New(),
		TxnRefId:  payload.TxnRefId,
		TxnTypeId: payload.TxnTypeId,
		TxnAmount: payload.TxnAmount,
		SpId:      payload.SpId,
	}

	// create new user
	createTxnResult := dbconn.DB.Create(&newTxn)
	if createTxnResult.Error != nil && strings.Contains(createTxnResult.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "duplicate key value violates unique"}))
	} else if createTxnResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Something bad happened"}))
	}

	response := toTxnHistoryResponse(newTxn)

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"txn_history": response}))
}
