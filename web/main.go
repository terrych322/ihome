package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "view")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFileStaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
