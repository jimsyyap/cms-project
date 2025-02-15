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
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    // PostgreSQL connection string
    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode

    // Connect to the database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // Run migrations for all models
    // err = DB.AutoMigrate(&models.User{}, &models.Post{}) // Add &models.Post{} here
    // if err != nil {
    //     panic("Failed to migrate database: " + err.Error())
    // }

    // Run migrations for all models
    err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Media{}) // Add &models.Media{}
    if err != nil {
        panic("Failed to migrate database: " + err.Error())
    }

    println("Database connected and migrated successfully!")
}
