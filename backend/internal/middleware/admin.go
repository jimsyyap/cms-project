package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// AdminMiddleware ensures the user is an admin
func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        if role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "only admins can perform this action"})
            c.Abort()
            return
        }

        c.Next()
    }
}
