package controller

import (
	"apiexample/src/model"
	"apiexample/src/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pay(context *gin.Context) {
	// var test string
	checkoutRequest := new(model.CheckoutRequest)
	stripeService := &service.StripeService{}
	// stripeService := service.NewStripeService("")

	// err := context.Bind(checkoutRequest)
	if err := context.Bind(checkoutRequest); err != nil {
		log.Printf("Error deserializing request body %s", err)
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	checkoutResponse, err := stripeService.GetUri(checkoutRequest)

	if err != nil {
		log.Printf("Error calling service %s", err)
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.JSON(http.StatusOK, checkoutResponse)
}
