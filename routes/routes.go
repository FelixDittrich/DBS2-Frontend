package routes

import (
	"de.stuttgart.hft/DBS2-Frontend/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(router *gin.Engine) {
	router.GET("/", controllers.RedirectToIndex)
	router.GET("/photos", controllers.OpenPhotos)
	router.POST("/photos", controllers.OpenPhotosByTypeId)
	router.GET("/deletephoto/:id", controllers.DeleteSinglePhoto)
	router.GET("/rolls", controllers.OpenRolls)
	router.POST("/createRoll", controllers.CreateRoll)
	router.POST("/uploadPhoto", controllers.UploadPhotos)
	router.GET("/roll/:id", controllers.OpenRollById)
	router.GET("/deleteroll/:id", controllers.DeleteRollAndPhotos)
	router.GET("/deletephotopool/:id", controllers.DeletePhotoFromPool)
	router.GET("/showphoto/:uuid", controllers.ShowPhoto)
}
