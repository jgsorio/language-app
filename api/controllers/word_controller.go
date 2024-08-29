package controllers

import (
	"language-app-api/services"

	"github.com/gin-gonic/gin"
)

type WordController struct{}

func (c WordController) Index(g *gin.Context) {
	words := services.IaService{}.SearchWords()
	
	g.JSON(200, words)
}