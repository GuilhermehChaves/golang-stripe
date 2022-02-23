package helper

import (
	"apiexample/src/constants"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func amountToCents(amount float32) int64 {
	return int64(amount * 100)
}

func getStripeCreateCheckoutData(successUrl, cancelUrl string, amount float32) url.Values {
	data := url.Values{}
	data.Set(constants.SuccessUrl, successUrl)
	data.Set(constants.CancelUrl, cancelUrl)
	data.Set(constants.Mode, "payment")
	data.Set(constants.PaymentMethodTypes, "card")
	data.Set(constants.LineItemsAmount, strconv.FormatInt(amountToCents(amount), 10))
	data.Set(constants.LineItemsCurrency, "BRL")
	data.Set(constants.LineItemsName, "TOTAL")
	data.Set(constants.LineItemsQuantity, "1")
	data.Set(constants.CaptureMethod, "manual")
	data.Set(constants.ClientReferenceId, "id_001")
	data.Set(constants.BillingAddressCollection, "required")

	return data
}

func CreateStripeRequest(url, successUrl, cancelUrl string, amount float32) (*http.Request, error) {
	data := getStripeCreateCheckoutData(successUrl, cancelUrl, amount)
	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	request.Header.Add(constants.Authorization, "Bearer stripe_secret_key")
	request.Header.Add(constants.ContentType, constants.FormUrlEnconded)
	request.Header.Add(constants.ContentLength, strconv.Itoa(len(data.Encode())))

	return request, nil
}
