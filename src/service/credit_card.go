package service

import (
	"apiexample/src/helper"
	"apiexample/src/interfaces"
	"apiexample/src/model"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type StripeService struct {
	Client interfaces.HttpClient
}

func NewStripeService(client interfaces.HttpClient) *StripeService {
	return &StripeService{
		Client: client,
	}
}

func (s *StripeService) GetUri(checkout *model.CheckoutRequest) (*model.CheckoutResponse, error) {
	request, err := helper.CreateStripeRequest(
		viper.GetString("STRIPE_URL"),
		checkout.SuccessUrl,
		checkout.CancelUrl,
		checkout.Amount,
	)

	defer s.Client.CloseIdleConnections()

	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	res, err := s.Client.Do(request)
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
