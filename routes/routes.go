package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manolors/full-api-example/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("coches", controller.GetCoches)
		v1.GET("coche", controller.GetCoche)
		v1.PUT("coche", controller.Nuevo)
	}
	return r
}
