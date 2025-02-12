package middleware

import (
    "github.com/jimsyyap/cms-project/internal/utils"
    "github.com/gin-gonic/gin"
    "net/http"
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

        // Verify the token
        claims, err := utils.VerifyJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Add the user ID and role to the context
        c.Set("userID", claims["user_id"])
        c.Set("role", claims["role"])

        // Continue to the next handler
        c.Next()
    }
}
