package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/models"
	"gorm.io/gorm"
)

type PhotoController struct {
	DB *gorm.DB
}

func (c *PhotoController) CreatePhoto(ctx *gin.Context) {

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	var photo models.Photo
	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(photo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set authenticated user ID as the creator of the photo
	photo.UserID = authenticatedUserID

	// Create photo
	err = c.DB.Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (c *PhotoController) GetPhotos(ctx *gin.Context) {

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	var photos []models.Photo
	c.DB.Where("user_id = ?", authenticatedUserID).Find(&photos)

	ctx.JSON(http.StatusOK, photos)
}

func (c *PhotoController) UpdatePhoto(ctx *gin.Context) {

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	// Get photo from database
	var photo models.Photo
	err := c.DB.First(&photo, ctx.Param("photoId")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check if the user is authorized to update the photo
	if photo.UserID != authenticatedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Update photo data
	var updatePhoto models.Photo
	err = ctx.ShouldBindJSON(&updatePhoto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(updatePhoto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update photo fields
	photo.Title = updatePhoto.Title
	photo.Caption = updatePhoto.Caption
	photo.PhotoURL = updatePhoto.PhotoURL

	// Save updated photo data
	err = c.DB.Save(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (c *PhotoController) DeletePhoto(ctx *gin.Context) {

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	// Get photo from database
	var photo models.Photo
	err := c.DB.First(&photo, ctx.Param("photoId")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check if the user is authorized to delete the photo
	if photo.UserID != authenticatedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete photo
	err = c.DB.Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}
