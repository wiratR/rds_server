package controllers

import (
	"errors"
	"fmt"
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
		TxnDetail: models.TxnDetail{
			TxnTypeId:   txn.TxnTypeId,
			TxnTypeName: getTxnTypeNameById(txn.TxnTypeId),
		},
		SpDetail: models.SpDetail{
			SpId:   txn.SpId,
			SpName: getSpNameById(txn.SpId),
		},
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

func toCreateTxn(txn models.TxnCreateInput, accountId uuid.UUID) int {

	if dbconn.DB == nil {
		return -1
	}

	newTxn := models.TxnHistory{
		ID:              uuid.New(),
		TxnRefId:        txn.TxnRefId,
		TxnTypeId:       txn.TxnTypeId,
		TxnAmount:       txn.TxnAmount,
		SpId:            txn.SpId,
		LocEntryId:      txn.LocEntryId,
		LocExitId:       txn.LocExitId,
		EquipmentNumber: txn.EquipmentNumber,
		AccountID:       accountId,
	}

	// create new user
	createTxnResult := dbconn.DB.Create(&newTxn)
	if createTxnResult.Error != nil && strings.Contains(createTxnResult.Error.Error(), "duplicate key value violates unique") {
		return -1
	} else if createTxnResult.Error != nil {
		return -1
	}

	return 0
}

func getTxnHistoriesByAccountId(accountId uuid.UUID) []models.TxnHistoryResponse {

	if dbconn.DB == nil {
		return nil
	}

	var txnHistories []models.TxnHistory
	var txnHistoryResponses []models.TxnHistoryResponse

	// Query the database for all TxnHistory records for a specific account
	if err := dbconn.DB.Where("account_id = ?", accountId).Find(&txnHistories).Error; err != nil {
		return nil
	}

	// Map TxnHistory records to TxnHistoryResponse
	for _, txn := range txnHistories {

		LineIdOfEntry := getLineIdByStationId(txn.LocEntryId)
		LineIdOfExit := getLineIdByStationId(txn.LocExitId)
		LineNameOfEntry := getLineNameByLineId(LineIdOfEntry)
		LineNameOfExit := getLineNameByLineId(LineIdOfExit)

		fmt.Println(LineIdOfEntry)
		fmt.Println(LineIdOfExit)

		response := models.TxnHistoryResponse{
			TxnRefId:  txn.TxnRefId,
			TxnAmount: txn.TxnAmount,
			// These details could be populated based on other queries or by using preloaded data
			TxnDetail: models.TxnDetail{
				TxnTypeId:   txn.TxnTypeId,
				TxnTypeName: getTxnTypeNameById(txn.TxnTypeId),
			},
			SpDetail: models.SpDetail{
				SpId:   txn.SpId,
				SpName: getSpNameById(txn.SpId),
			},
			LocEntryDetail: models.LocDetail{
				LocId:   txn.LocEntryId,
				LocName: getStationNameById(txn.LocEntryId),
				LineDetail: models.LineDetail{
					LineId:   LineIdOfEntry,
					LineName: LineNameOfEntry,
				},
			},
			LocExitDetail: models.LocDetail{
				LocId:   txn.LocExitId,
				LocName: getStationNameById(txn.LocExitId),
				LineDetail: models.LineDetail{
					LineId:   LineIdOfExit,
					LineName: LineNameOfExit,
				},
			},
			EquipmentNumber: txn.EquipmentNumber,
			CreatedAt:       txn.CreatedAt,
			UpdatedAt:       txn.UpdatedAt,
		}

		txnHistoryResponses = append(txnHistoryResponses, response)
	}

	// Return the slice of TxnHistoryResponse records
	return txnHistoryResponses

}

// GetTxnHistory retrieves a transaction history by account ID
// @Description retrieves a transaction history by account ID
// @Summary retrieves a transaction history by account ID
// @Tags Txnhistories
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {array} models.TxnHistoryResponse
// @Router /Txnhistories/{accountId} [get]
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
// @Tags Txnhistories
// @Accept json
// @Produce json
// @Param Payload body models.TxnCreateInput true "Txn Create Input Data"
// @Success 201 {array} models.TxnHistoryResponse "Ok"
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
