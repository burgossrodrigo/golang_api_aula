package main

import (
	"golang-api/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eg := r.Group("/v1")

	{
		eg.GET("/horas", horas)
		eg.POST("/horas", horasComEducação)
	}

	r.Run(":8080")
}

func horas(c *gin.Context) {
	c.JSON(200, gin.H{
		"horas": time.Now().Format("15:04:05"),
	})
}

func horasComEducação(c *gin.Context) {

	var mensagemModel models.HorasMensagem

	if err := c.ShouldBindJSON(&mensagemModel); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	if c.Param("mensagem") != "por_favor" {
		c.JSON(200, gin.H{
			"horas":    time.Now().Format("15:04:05"),
			"mensagem": "Parabéns! Você é educado!",
		})
	} else {
		c.JSON(400, gin.H{
			"error": "palavra mágica?",
		})
	}
}
