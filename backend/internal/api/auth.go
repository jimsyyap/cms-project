package api

import (
    "github.com/jimsyyap/cms-project/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context) {
    var req struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=8"`
        Role     string `json:"role" binding:"required"`
    }

    // Bind the request body to the struct
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Register the user
    user, err := services.RegisterUser(req.Username, req.Email, req.Password, req.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the created user (excluding sensitive data)
    c.JSON(http.StatusCreated, gin.H{
        "id":       user.ID,
        "username": user.Username,
        "email":    user.Email,
        "role":     user.Role,
    })
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    // Bind the request body to the struct
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Authenticate the user
    token, err := services.LoginUser(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Return the JWT token
    c.JSON(http.StatusOK, gin.H{"token": token})
}
