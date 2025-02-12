// backend/main.go
package main

import (
    "github.com/jimsyyap/cms-project/internal/api"
    "github.com/jimsyyap/cms-project/internal/config"
    "github.com/jimsyyap/cms-project/internal/db"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Initialize database
    db.InitDB()

    // Create Gin router
    r := gin.Default()

    // Register API routes
    api.RegisterRoutes(r)

    // Start the server
    r.Run(":8080")
}
