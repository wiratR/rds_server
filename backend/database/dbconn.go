package database

import (
	"fmt"
	"log"
	"os"

	"github.com/wiratR/go-orm-jwt/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	var err error

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	fmt.Printf("Database hostname = %s\n", config.DBHost)
	fmt.Printf("Database port = %d\n", config.DBPort)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
	// 	config.DBHost,
	// 	config.DBUserName,
	// 	config.DBUserPassword,
	// 	config.DBName,
	// 	config.DBPort)

	// Build connection string
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=Asia/Bangkok"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Database connection established")
	}

	log.Println("🚀 Connected Successfully to the Database")
}
