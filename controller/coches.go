package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manolors/full-api-example/model"
)

func GetCoches(c *gin.Context) {
	var coches model.Coches
	coches.Get()
	c.JSON(http.StatusOK, gin.H{"data": coches})
}

func GetCoche(c *gin.Context) {
	var coche model.Coche
	err := c.BindQuery(&coche)
	if err != nil {
		log.Println("Error al bindear datos", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not implemented"})
	}
}

func Nuevo(c *gin.Context) {
	var coche model.Coche
	err := c.BindJSON(&coche)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect fields"})
		log.Println("Error al bindear datos", err)
		return
	}
	coche.Nuevo()
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not implemented"})
}
