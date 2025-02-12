package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username     string `gorm:"size:50;unique;not null"`
    Email        string `gorm:"size:100;unique;not null"`
    PasswordHash string `gorm:"not null"`
    Role         string `gorm:"size:20;not null;default:'member'"`
    CreatedAt    time.Time
}
