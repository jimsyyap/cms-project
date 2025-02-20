package main

import (
    "github.com/jimsyyap/cms-project/internal/api"
    "github.com/jimsyyap/cms-project/internal/config"
    "github.com/jimsyyap/cms-project/internal/db"
    "github.com/jimsyyap/cms-project/internal/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Initialize database and run migrations
    db.InitDB()

    // Create Gin router
    r := gin.Default()

    // Apply global middleware
    r.Use(middleware.LoggingMiddleware())

    // Public routes (no authentication required)
    r.POST("/register", api.RegisterHandler)
    r.POST("/login", api.LoginHandler)

    // Protected routes (require authentication)
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        // User management routes
        protected.GET("/users", api.GetAllUsersHandler)
        protected.GET("/users/:id", api.GetUserByIDHandler)

        // Admin-only routes
        protected.PUT("/users/:id/role", middleware.AdminMiddleware(), api.UpdateUserRoleHandler)
        protected.DELETE("/users/:id", middleware.AdminMiddleware(), api.DeleteUserHandler)
    }

    // Start the server
    r.Run(":8080")
}
