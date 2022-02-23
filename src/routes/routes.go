package routes

import (
	"apiexample/src/factory"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	route := gin.New()
	creditCardController := factory.NewCreditCardController()

	route.POST("/card/pay", creditCardController.Pay)
	return route
}
