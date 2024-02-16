package main

import (
	"github.com/adasarpan404/change/controllers"
	"github.com/adasarpan404/change/environment"
	"github.com/adasarpan404/change/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(controllers.Authenticate())
	routes.UserRoutes(router)
	router.Run(":" + environment.PORT)
}
