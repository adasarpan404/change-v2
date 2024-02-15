package routes

import (
	"github.com/adasarpan404/change/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/user/me", controllers.GetUser())
}
