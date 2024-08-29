package controllers

import (
	"language-app-api/models"

	"github.com/gin-gonic/gin"
)

type HomeController struct {}

func (c HomeController) Index(g *gin.Context) {
	g.JSON(200, models.Message{Body: "Hello, World!"})
}