package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/controllers"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	userController := &controllers.UserController{DB: db}
	photoController := &controllers.PhotoController{DB: db}

	r.POST("/users/register", userController.RegisterUser)
	r.GET("/users/login", userController.LoginUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)

	r.POST("/photos", middlewares.AuthMiddleware(), photoController.CreatePhoto)
	r.GET("/photos", middlewares.AuthMiddleware(), photoController.GetPhotos)
	r.PUT("/photos/:photoId", middlewares.AuthMiddleware(), photoController.UpdatePhoto)
	r.DELETE("/photos/:photoId", middlewares.AuthMiddleware(), photoController.DeletePhoto)
}
