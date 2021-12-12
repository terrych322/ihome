package main

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func main() {
	router := gin.Default()
	router.Static("/home", "view")
	router.GET("/api/v1.0/session", controller.GetSession)
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
