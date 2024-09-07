package main

import (
	"fmt"
	"log"
	"os"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/wiratR/rds_server/config"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	"github.com/wiratR/rds_server/routes"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/wiratR/rds_server/docs"
)

func setupRoutes(app *fiber.App) {

	// Routes
	app.Get("/", HealthCheck)

	// api group
	api := app.Group("/api")
	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// connect Auth routes
	routes.AuthRoute(api.Group("/auth"))
	// connect User routes
	routes.UserRoute(api.Group("/users"))
	// connect Payment routes
	routes.PaymentRoute(api.Group("/payment"))
	// connect Account routes
	routes.AccountRoute(api.Group("/accounts"))
	// connect TxnHistory routes
	routes.AccountRoute(api.Group("/txnhistories"))
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /api
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	fmt.Println("main start")
	fmt.Printf("get Environment = %s\n", config.Environment)
	fmt.Printf("get Version = %d\n", config.Version)

	dbconn.InitDatabase()

	// Auto-migrate the schema
	dbconn.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	// dbconn.DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// Perform database migrations
	if err := dbconn.DB.AutoMigrate(&models.User{}, &models.Account{}); err != nil {
		log.Fatalln("failed to migrate database:", err)
		os.Exit(1)
	}
	// AutoMigrate the schema for parameter
	err = dbconn.DB.AutoMigrate(&models.CardType{}, &models.MediaType{}, &models.Station{}, &models.Line{}, &models.ServiceProvider{}, &models.Fare{}, &models.TxnType{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	// AutoMigrate the schema for transaction history
	err = dbconn.DB.AutoMigrate(&models.TxnHistory{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// create seed parameter
	seed(dbconn.DB)

	app := fiber.New()
	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(etag.New())

	// app.Use(cors.New(cors.Config{
	// 	AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	// 	AllowOrigins:     "*",
	// 	AllowCredentials: true,
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	// }))

	// app.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"http://example.com", "http://another-allowed-origin.com"},
	//     AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	//     AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	//     MaxAge:           12 * time.Hour,
	// }))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// setup routes
	setupRoutes(app) // new

	// Listen on server 8000 and catch error if any
	err = app.Listen(":" + config.Port)

	// handle error
	if err != nil {
		panic(err)
	}

}

func seed(db *gorm.DB) {
	var count int64

	// Check and seed card types
	db.Model(&models.CardType{}).Count(&count)
	if count == 0 {
		cardTypes := []models.CardType{
			{ID: uuid.New(), CardId: 1, ShortName: "ADL", Description: "Adult Card", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), CardId: 2, ShortName: "STU", Description: "Student Card", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&cardTypes).Error; err != nil {
			log.Fatalf("failed to seed card types: %v", err)
		}
	} else {
		log.Println("Card types already seeded")
	}

	// Check and seed media types
	db.Model(&models.MediaType{}).Count(&count)
	if count == 0 {
		mediaTypes := []models.MediaType{
			{ID: uuid.New(), MediaTypeId: 1, ShortName: "CSC", Description: "Card", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), MediaTypeId: 2, ShortName: "CST", Description: "Token", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&mediaTypes).Error; err != nil {
			log.Fatalf("failed to seed media types: %v", err)
		}
	} else {
		log.Println("Media types already seeded")
	}

	// Check and seed stations
	db.Model(&models.Station{}).Count(&count)
	if count == 0 {
		stations := []models.Station{
			{ID: uuid.New(), StationId: 1, ShortName: "STA1", Description: "Station 1", IsCrossLine: false, LineId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), StationId: 2, ShortName: "STA2", Description: "Station 2", IsCrossLine: true, LineId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), StationId: 99, ShortName: "RTVS", Description: "RTV service", IsCrossLine: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&stations).Error; err != nil {
			log.Fatalf("failed to seed stations: %v", err)
		}
	} else {
		log.Println("Stations already seeded")
	}

	// Check and seed lines
	db.Model(&models.Line{}).Count(&count)
	if count == 0 {
		lines := []models.Line{
			{ID: uuid.New(), LineId: 0, ShortName: "L0", Description: "Line 0", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), LineId: 1, ShortName: "L1", Description: "Line 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), LineId: 2, ShortName: "L2", Description: "Line 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&lines).Error; err != nil {
			log.Fatalf("failed to seed lines: %v", err)
		}
	} else {
		log.Println("Lines already seeded")
	}

	// Check and seed service providers
	db.Model(&models.ServiceProvider{}).Count(&count)
	if count == 0 {
		serviceProviders := []models.ServiceProvider{
			{ID: uuid.New(), ServiceProviderId: 1, ShortName: "SP1", Description: "Service Provider 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), ServiceProviderId: 2, ShortName: "SP2", Description: "Service Provider 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), ServiceProviderId: 9, ShortName: "CCS", Description: "Center Clearing Server", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&serviceProviders).Error; err != nil {
			log.Fatalf("failed to seed service providers: %v", err)
		}
	} else {
		log.Println("Service providers already seeded")
	}

	// Check and seed fares
	db.Model(&models.Fare{}).Count(&count)
	if count == 0 {
		fares := []models.Fare{
			{ID: uuid.New(), SpId: 1, LineId: 1, StationId: 1, CardTypeId: 1, MediaTypeId: 1, Amount: 100, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), SpId: 1, LineId: 1, StationId: 2, CardTypeId: 1, MediaTypeId: 1, Amount: 200, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), SpId: 2, LineId: 2, StationId: 1, CardTypeId: 2, MediaTypeId: 1, Amount: 100, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), SpId: 2, LineId: 2, StationId: 2, CardTypeId: 2, MediaTypeId: 1, Amount: 200, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&fares).Error; err != nil {
			log.Fatalf("failed to seed fares: %v", err)
		}
	} else {
		log.Println("Fares already seeded")
	}

	// Check and seed Transaction type
	db.Model(&models.TxnType{}).Count(&count)
	if count == 0 {
		txnTypes := []models.TxnType{
			{ID: uuid.New(), TxnTypeId: 1, Description: "Binding", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 2, Description: "Check In", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 3, Description: "Check Out", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 4, Description: "Upgarde", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 5, Description: "Block", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 6, Description: "Un Block", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 7, Description: "Add Value", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: uuid.New(), TxnTypeId: 8, Description: "Auto Add Value", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Create(&txnTypes).Error; err != nil {
			log.Fatalf("failed to seed Transaction type: %v", err)
		}
	} else {
		log.Println("Transaction type already seeded")
	}

	log.Println("Database seeded successfully")
}

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
