package controller

import (
	"apiexample/src/interfaces"
	"apiexample/src/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditCardController struct {
	service interfaces.CreditCardService
}

func NewCreditCardController(service interfaces.CreditCardService) *CreditCardController {
	return &CreditCardController{
		service: service,
	}
}

func (c *CreditCardController) Pay(context *gin.Context) {
	// var test string
	checkoutRequest := new(model.CheckoutRequest)

	// err := context.Bind(checkoutRequest)
	if err := context.Bind(checkoutRequest); err != nil {
		log.Printf("Error deserializing request body %s", err)
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	checkoutResponse, err := c.service.GetUri(checkoutRequest)

	if err != nil {
		log.Printf("Error calling service %s", err)
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.JSON(http.StatusOK, checkoutResponse)
}
