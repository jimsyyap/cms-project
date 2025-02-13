package services

import (
    "github.com/jimsyyap/cms-project/internal/models"
    "github.com/jimsyyap/cms-project/internal/db"
    // "errors"
)

// CreatePost creates a new post
func CreatePost(title, content string, authorID uint) (*models.Post, error) {
    post := models.Post{
        Title:    title,
        Content:  content,
        AuthorID: authorID,
    }

    result := db.DB.Create(&post)
    if result.Error != nil {
        return nil, result.Error
    }

    return &post, nil
}

// GetPostByID retrieves a post by its ID
func GetPostByID(id uint) (*models.Post, error) {
    var post models.Post
    result := db.DB.First(&post, id)
    if result.Error != nil {
        return nil, result.Error
    }

    return &post, nil
}

// UpdatePost updates an existing post
func UpdatePost(id uint, title, content string) (*models.Post, error) {
    var post models.Post
    result := db.DB.First(&post, id)
    if result.Error != nil {
        return nil, result.Error
    }

    post.Title = title
    post.Content = content

    result = db.DB.Save(&post)
    if result.Error != nil {
        return nil, result.Error
    }

    return &post, nil
}

// DeletePost deletes a post by its ID
func DeletePost(id uint) error {
    result := db.DB.Delete(&models.Post{}, id)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

// GetAllPosts retrieves all posts
func GetAllPosts() ([]models.Post, error) {
    var posts []models.Post
    result := db.DB.Find(&posts)
    if result.Error != nil {
        return nil, result.Error
    }

    return posts, nil
}
