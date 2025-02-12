package db

import (
    "github.com/jimsyyap/cms-project/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    // PostgreSQL connection string
    dsn := "host=localhost user=jim password=definitionemotionexperience dbname=cmsdb port=5432 sslmode=disable"

    // Connect to the database
    var err error
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
