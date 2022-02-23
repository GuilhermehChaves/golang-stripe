package interfaces

import "apiexample/src/model"

type CreditCardService interface {
	GetUri(checkout *model.CheckoutRequest) (*model.CheckoutResponse, error)
}
