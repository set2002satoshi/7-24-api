package router

import (
	// "log"
	

	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/7-24-api/controller"
)

func SetRouter() {

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", controller.FindComment)
		v1.POST("/", controller.CreateComment)
		v1.DELETE("/", controller.DeleteTarget)
	}

	router.Run(":8080")
}