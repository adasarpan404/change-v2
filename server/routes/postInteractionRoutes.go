package routes

import (
	"github.com/adasarpan404/change/controllers"
	"github.com/gin-gonic/gin"
)

func PostInteractionRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/like", controllers.GetLikes())
	incomingRoutes.POST("/like", controllers.CreateLike())
	incomingRoutes.PUT("/like/:id", controllers.UpdateLike())
	incomingRoutes.DELETE("/like/:id", controllers.DeleteLike())

	incomingRoutes.GET("/comment", controllers.GetComments())
	incomingRoutes.POST("/comment", controllers.CreateComment())
	incomingRoutes.PUT("/comment/:id", controllers.UpdateComment())
	incomingRoutes.DELETE("/comment/:id", controllers.DeleteComment())
}
