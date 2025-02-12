package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

// LoggingMiddleware logs the request method, path, and response time
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // Process the request
        c.Next()

        // Log the request details
        log.Printf(
            "Method: %s | Path: %s | Status: %d | Duration: %s",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start),
        )
    }
}
