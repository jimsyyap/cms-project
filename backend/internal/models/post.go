package models

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title    string `gorm:"size:255;not null"`
    Content  string `gorm:"type:text;not null"`
    AuthorID uint   `gorm:"not null"` // Foreign key to the User model
}
