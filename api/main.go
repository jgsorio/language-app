package main

import (
	"language-app-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/", controllers.HomeController{}.Index)
	r.GET("/words", controllers.WordController{}.Index)

	r.Run(":8080")
}
