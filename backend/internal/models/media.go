package models

import (
    "gorm.io/gorm"
)

type Media struct {
    gorm Model
    FileName string `gorm:"size:255;not null"` // name of upload file
    FileURL   string `gorm:"not null"`          // URL or path to the file
    UploadedBy uint  `gorm:"not null"`          // ID of the user who uploaded the file
}
