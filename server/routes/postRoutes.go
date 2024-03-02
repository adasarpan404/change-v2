package routes

import (
	"github.com/adasarpan404/change/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/post", controllers.GetPost())
	incomingRoutes.POST("/post", controllers.CreatePost())
	incomingRoutes.PUT("/post/:id", controllers.UpdatePost())
	incomingRoutes.DELETE("/post/:id", controllers.DeletePost())
}
