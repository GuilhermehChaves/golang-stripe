package service

import (
	"apiexample/src/helper"
	"apiexample/src/model"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StripeService struct {
	Url string
}

func NewStripeService(url string) *StripeService {
	return &StripeService{
		Url: url,
	}
}

func (s *StripeService) GetUri(checkout *model.CheckoutRequest) (*model.CheckoutResponse, error) {
	url := "https://api.stripe.com/v1/checkout/sessions"

	client := new(http.Client) // &http.Client{}
	request, err := helper.CreateStripeRequest(url, checkout.SuccessUrl, checkout.CancelUrl, checkout.Amount)

	defer client.CloseIdleConnections()

	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	res, err := client.Do(request)
	if err != nil {
		log.Printf("Error perfoming request to Stripe: %s", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Erro during request to Stripe [status code: %d]", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("Error status code: %d", res.StatusCode))
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error %s", err)
		return nil, err
	}

	defer res.Body.Close()
	checkoutResponse := new(model.CheckoutResponse)

	if err := helper.ResponseToStruct(body, checkoutResponse); err != nil {
		log.Printf("Error converting response to struct %s", err)
		return nil, err
	}

	return checkoutResponse, nil
}
