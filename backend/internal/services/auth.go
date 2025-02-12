package services

import (
    "github.com/jimsyyap/cms-project/internal/utils"
    "github.com/jimsyyap/cms-project/internal/models"
    "github.com/jimsyyap/cms-project/internal/db"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user
func RegisterUser(username, email, password, role string) (*models.User, error) {
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Create a new user
    user := models.User{
        Username:     username,
        Email:        email,
        PasswordHash: string(hashedPassword),
        Role:         role,
    }

    // Insert the user into the database
    result := db.DB.Create(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    return &user, nil
}

// LoginUser authenticates a user and returns a JWT
func LoginUser(email, password string) (string, error) {
    // Fetch the user from the database
    var user models.User
    result := db.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return "", errors.New("invalid email or password")
    }

    // Compare the provided password with the stored hash
    err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    // Generate a JWT
    token, err := utils.GenerateJWT(user.ID, user.Role) // Pass user.ID directly (no cast needed)
    if err != nil {
        return "", err
    }

    return token, nil
}
