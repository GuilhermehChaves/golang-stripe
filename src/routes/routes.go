package routes

import (
	"apiexample/src/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	route := gin.New()

	route.POST("/card/pay", controller.Pay)
	return route
}
