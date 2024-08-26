package controllers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	"gorm.io/gorm"
)

// Get all CardType records
func GetCardTypeAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	var cardTypes []models.CardType

	if err := dbconn.DB.Find(&cardTypes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, cardTypes))
}

// Get a single CardType by ID
func GetCardType(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var cardType models.CardType

	if err := dbconn.DB.First(&cardType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, cardType))
}

// Update a CardType by ID
func UpdateCardType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var cardType models.CardType

	if err := dbconn.DB.First(&cardType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&cardType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&cardType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, cardType))
}

// Delete a CardType by ID
func DeleteCardType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var cardType models.CardType

	if err := dbconn.DB.First(&cardType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&cardType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all MediaType records
func GetMediaTypeAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var mediaTypes []models.CardType

	if err := dbconn.DB.Find(&mediaTypes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, mediaTypes))
}

// Get a single MediaType by ID
func GetMediaType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var mediaType models.MediaType

	if err := dbconn.DB.First(&mediaType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, mediaType))
}

// Update a MediaType by ID
func UpdateMediaType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var mediaType models.MediaType

	if err := dbconn.DB.First(&mediaType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&mediaType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&mediaType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, mediaType))
}

// Delete a MediaType by ID
func DeleteMediaType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var mediaType models.MediaType

	if err := dbconn.DB.First(&mediaType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&mediaType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all Station records
func GetStationAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var stations []models.Station

	if err := dbconn.DB.Find(&stations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, stations))
}

// Get a single Station by ID
func GetStation(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var station models.Station

	if err := dbconn.DB.First(&station, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, station))
}

// Update a Station by ID
func UpdateStation(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var station models.Station

	if err := dbconn.DB.First(&station, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&station); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&station).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, station))
}

// Delete a Station by ID
func DeleteStation(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var station models.Station

	if err := dbconn.DB.First(&station, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&station).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all Line records
func GetLineAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var lines []models.Line

	if err := dbconn.DB.Find(&lines).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, lines))
}

// Get a single Line by ID
func GetLine(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var line models.Line

	if err := dbconn.DB.First(&line, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, line))
}

// Update a Line by ID
func UpdateLine(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var line models.Line

	if err := dbconn.DB.First(&line, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&line); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&line).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, line))
}

// Delete a Line by ID
func DeleteLine(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var line models.Line

	if err := dbconn.DB.First(&line, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&line).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all ServiceProviders records
func GetServiceProviderAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var serviceProviders []models.ServiceProvider

	if err := dbconn.DB.Find(&serviceProviders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, serviceProviders))
}

// Get a single ServiceProvider by ID
func GetServiceProvider(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var serviceProvider models.ServiceProvider

	if err := dbconn.DB.First(&serviceProvider, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, serviceProvider))
}

// Update a ServiceProviders by ID
func UpdateServiceProvider(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var serviceProvider models.ServiceProvider

	if err := dbconn.DB.First(&serviceProvider, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&serviceProvider); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&serviceProvider).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, serviceProvider))
}

// Delete a ServiceProvider by ID
func DeleteServiceProvider(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var serviceProvider models.ServiceProvider

	if err := dbconn.DB.First(&serviceProvider, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&serviceProvider).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all Fares records
func GetFareAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var fares []models.Fare

	if err := dbconn.DB.Find(&fares).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fares))
}

// Get a single Fare by ID
func GetFare(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var fare models.Fare

	if err := dbconn.DB.First(&fare, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fare))
}

// Update a Fare by ID
func UpdateFare(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var fare models.Fare

	if err := dbconn.DB.First(&fare, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&fare); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&fare).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fare))
}

// Delete a Fare by ID
func DeleteFare(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var fare models.Fare

	if err := dbconn.DB.First(&fare, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&fare).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

// Get all TxnType records
func GetTxnTypeAll(c *fiber.Ctx) error {
	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false
	var txnTypes []models.TxnType

	if err := dbconn.DB.Find(&txnTypes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, txnTypes))
}

// Get a single TxnType by ID
func GetTxnType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}
	var isSuccess bool = false

	id := c.Params("id")
	var txnType models.TxnType

	if err := dbconn.DB.First(&txnType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, txnType))
}

// Update a TxnType by ID
func UpdateTxnType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var txnType models.TxnType

	if err := dbconn.DB.First(&txnType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := c.BodyParser(&txnType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Save(&txnType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, txnType))
}

// Delete a TxnType by ID
func DeleteTxnType(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false
	id := c.Params("id")
	var txnType models.TxnType

	if err := dbconn.DB.First(&txnType, "id = ?", uuid.MustParse(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "the given ID is not found"}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	if err := dbconn.DB.Delete(&txnType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse(isSuccess, fiber.Map{"error": err.Error()}))
	}

	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": "Record deleted"}))
}

func getStationNameById(stationId int) string {
	var station models.Station
	result := dbconn.DB.First(&station, "StationId = ?", stationId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Station with ID %d not found", stationId)
			return ""
		}
		log.Printf("Error retrieving station: %v", result.Error)
		return ""
	}

	return station.Description
}

func getLineIdByStationId(stationId int) int {
	var station models.Station
	result := dbconn.DB.First(&station, "StationId = ?", stationId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Station with ID %d not found", stationId)
			return -1
		}
		log.Printf("Error retrieving station: %v", result.Error)
		return -1
	}

	return station.LineId
}

func getLineNameByLineId(lineId int) string {
	var line models.Line
	result := dbconn.DB.First(&line, "LineId = ?", lineId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Station with ID %d not found", lineId)
			return ""
		}
		log.Printf("Error retrieving station: %v", result.Error)
		return ""
	}

	return line.Description
}
