package db

import (
    "github.com/jimsyyap/cms-project/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "os"
)

var DB *gorm.DB

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        panic("Failed to load .env file: " + err.Error())
    }

    // Get database configuration from environment variables
    dbHost, exists := os.LookupEnv("DB_HOST")
    if !exists {
        panic("DB_HOST environment variable is not set")
    }

    dbPort, exists := os.LookupEnv("DB_PORT")
    if !exists {
        panic("DB_PORT environment variable is not set")
    }

    dbUser, exists := os.LookupEnv("DB_USER")
    if !exists {
        panic("DB_USER environment variable is not set")
    }

    dbPassword, exists := os.LookupEnv("DB_PASSWORD")
    if !exists {
        panic("DB_PASSWORD environment variable is not set")
    }

    dbName, exists := os.LookupEnv("DB_NAME")
    if !exists {
        panic("DB_NAME environment variable is not set")
    }

    dbSSLMode, exists := os.LookupEnv("DB_SSLMODE")
    if !exists {
        dbSSLMode = "disable" // Default to "disable" if not set
    }

    // PostgreSQL connection string
    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode

    // Connect to the database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // Run migrations
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        panic("Failed to migrate database: " + err.Error())
    }

    println("Database connected and migrated successfully!")
}
