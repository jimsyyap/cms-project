package api

import (
    "github.com/jimsyyap/cms-project/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// CreatePostHandler handles creating a new post
func CreatePostHandler(c *gin.Context) {
    var req struct {
        Title   string `json:"title" binding:"required"`
        Content string `json:"content" binding:"required"`
    }

    // Bind the request body to the struct
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get the author ID from the JWT token
    authorID := c.GetUint("userID")

    // Create the post
    post, err := services.CreatePost(req.Title, req.Content, authorID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, post)
}

// GetPostHandler retrieves a post by its ID
func GetPostHandler(c *gin.Context) {
    // Get the post ID from the URL parameter
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
        return
    }

    // Fetch the post
    post, err := services.GetPostByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
        return
    }

    c.JSON(http.StatusOK, post)
}

// UpdatePostHandler updates an existing post
func UpdatePostHandler(c *gin.Context) {
    // Get the post ID from the URL parameter
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
        return
    }

    var req struct {
        Title   string `json:"title" binding:"required"`
        Content string `json:"content" binding:"required"`
    }

    // Bind the request body to the struct
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the post
    post, err := services.UpdatePost(uint(id), req.Title, req.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

// DeletePostHandler deletes a post by its ID
func DeletePostHandler(c *gin.Context) {
    // Get the post ID from the URL parameter
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
        return
    }

    // Delete the post
    err = services.DeletePost(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}

// GetAllPostsHandler retrieves all posts
func GetAllPostsHandler(c *gin.Context) {
    posts, err := services.GetAllPosts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, posts)
}
