package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/helpers"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserController struct {
	DB *gorm.DB
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user with the same email already exists
	var existingUser models.User
	c.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User with the same email already exists"})
		return
	}

	// Create user
	result := c.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var login models.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user with the provided email exists
	var existingUser models.User
	existingUser.Email = login.Email
	existingUser.Password = login.Password
	c.DB.Where("email = ?", login.Email).First(&existingUser)
	if existingUser.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate token
	token, err := helpers.GenerateToken(existingUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	// Get user from database
	var user models.User
	err := c.DB.First(&user, ctx.Param("userId")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if the user is authorized to update the profile
	if user.ID != authenticatedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Update user data
	var updateUser models.User
	err = ctx.ShouldBindJSON(&updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user fields
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	// Save updated user data
	err = c.DB.Save(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	authenticatedUserID := userID.(uint)

	// Get user from database
	var user models.User
	err := c.DB.First(&user, ctx.Param("userId")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if the user is authorized to delete the profile
	if user.ID != authenticatedUserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete user
	err = c.DB.Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
