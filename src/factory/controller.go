package factory

import (
	"apiexample/src/controller"
	"apiexample/src/service"
	"net/http"
)

var httpClient *http.Client = new(http.Client)

func NewCreditCardController() *controller.CreditCardController {
	stripeService := service.NewStripeService(httpClient)
	return controller.NewCreditCardController(stripeService)
}
