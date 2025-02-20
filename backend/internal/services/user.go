package services

import (
    "github.com/jimsyyap/cms-project/internal/models"
    "github.com/jimsyyap/cms-project/internal/db"
    "errors"
)

// GetAllUsers fetches all users
func GetAllUsers() ([]models.User, error) {
    var users []models.User
    result := db.DB.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }

    // Exclude sensitive data (e.g., PasswordHash)
    for i := range users {
        users[i].PasswordHash = ""
    }

    return users, nil
}

// GetUserByID fetches a user by their ID
func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    result := db.DB.First(&user, id)
    if result.Error != nil {
        return nil, errors.New("user not found")
    }

    // Exclude sensitive data
    user.PasswordHash = ""

    return &user, nil
}

// UpdateUserRole updates the role of a user
func UpdateUserRole(id uint, role string) (*models.User, error) {
    var user models.User
    result := db.DB.First(&user, id)
    if result.Error != nil {
        return nil, errors.New("user not found")
    }

    // Update the role
    user.Role = role
    result = db.DB.Save(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    // Exclude sensitive data
    user.PasswordHash = ""

    return &user, nil
}

// DeleteUser deletes a user by their ID
func DeleteUser(id uint) error {
    result := db.DB.Delete(&models.User{}, id)
    if result.Error != nil {
        return result.Error
    }

    return nil
}
