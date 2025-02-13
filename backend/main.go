// backend/main.go
package main

import (
    "net/http"
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
    r.Use(middleware.LoggingMiddleware()) // Add logging middleware
    r.Use(middleware.CORSMiddleware()) // Example: Add CORS middleware
    // r.Use(middleware.RateLimitMiddleware()) Example: Add rate-limiting middleware

    // Public routes (no authentication required)
    r.POST("/register", api.RegisterHandler)
    r.POST("/login", api.LoginHandler)

    // Protected routes (require authentication)
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware()) // Apply the AuthMiddleware
    {
        protected.GET("/protected", func(c *gin.Context) {
            userID := c.GetInt("userID")
            role := c.GetString("role")
            c.JSON(http.StatusOK, gin.H{"userID": userID, "role": role})
        })
        // Post management routes
        protected.POST("/posts", api.CreatePostHandler)
        protected.GET("/posts/:id", api.GetPostHandler)
        protected.PUT("/posts/:id", api.UpdatePostHandler)
        protected.DELETE("/posts/:id", api.DeletePostHandler)
        protected.GET("/posts", api.GetAllPostsHandler)
    }

    // Start the server
    r.Run(":8080")
}
