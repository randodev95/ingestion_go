package routes

import (
	"ingestion_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user/newregistration", controllers.CreateUser())
	router.GET("/user/findUser/:UserId", controllers.GetUserDetails())
	router.POST("user/editUser/:UserId", controllers.EditUserDetails())
	router.DELETE("user/deleteUser/:UserId", controllers.DeleteUser())
}

func PingRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping())
}
