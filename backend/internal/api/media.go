package api

import (
    "github.com/jimsyyap/cms-project/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

// UploadMediaHandler handles file uploads
func UploadMediaHandler(c *gin.Context) {
    // Get the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
        return
    }

    // Get the user ID from the JWT token
    userID := c.GetUint("userID")

    // Upload the file
    media, err := services.UploadMedia(file, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, media)
}
