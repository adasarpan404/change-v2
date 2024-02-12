package main

import (
	"github.com/adasarpan404/change/environment"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Run(":" + environment.PORT)
}
