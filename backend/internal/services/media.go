package services

import (
    "github.com/jimsyyap/cms-project/internal/models"
    "github.com/jimsyyap/cms-project/internal/db"
    "errors"
    "io" // Add this import
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// UploadMedia handles file uploads and saves metadata to the database
func UploadMedia(file *multipart.FileHeader, uploadedBy uint) (*models.Media, error) {
    // Generate a unique file name
    ext := filepath.Ext(file.Filename)
    fileName := strings.TrimSuffix(file.Filename, ext) + "-" + time.Now().Format("20060102150405") + ext

    // Define the file path
    filePath := filepath.Join("uploads", fileName)

    // Save the file to the uploads directory
    if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
        return nil, errors.New("failed to create uploads directory")
    }
    if err := saveUploadedFile(file, filePath); err != nil {
        return nil, errors.New("failed to save file")
    }

    // Save metadata to the database
    media := models.Media{
        FileName:   fileName,
        FileURL:    "/uploads/" + fileName, // URL to access the file
        UploadedBy: uploadedBy,
    }

    result := db.DB.Create(&media)
    if result.Error != nil {
        return nil, result.Error
    }

    return &media, nil
}

// saveUploadedFile saves the uploaded file to the specified path
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    if _, err := io.Copy(out, src); err != nil { // io.Copy requires the "io" package
        return err
    }

    return nil
}
