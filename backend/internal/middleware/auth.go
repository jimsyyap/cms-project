package middleware

import (
    "github.com/jimsyyap/cms-project/internal/utils"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
)

// AuthMiddleware verifies the JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the token from the Authorization header
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        log.Println("Received token:", token) // Log the token for debugging

        // Verify the token
        claims, err := utils.VerifyJWT(token)
        if err != nil {
            log.Println("Token verification failed:", err) // Log the error for debugging
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Add the user ID and role to the context
        c.Set("userID", claims["user_id"])
        c.Set("role", claims["role"])

        c.Next()
    }
}
