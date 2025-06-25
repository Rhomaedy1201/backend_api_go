package controllers

import (
	"backend-api/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// type validation post input
type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// function get error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	//return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "List data posts",
		"data":    posts,
	})
}

// store a post
func StorePost(c *gin.Context) {
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}
	// create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	//return json response
	c.JSON(201, gin.H{
		"success": true,
		"message": "Post created successfully",
		"data":    post,
	})
}

// get posts by id
func FindPostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By Id : " + c.Param("id"),
		"data":    post,
	})
}

// update posts
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//validete input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	//update post
	models.DB.Model(&post).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Posts updated successfully",
		"data":    post,
	})
}
