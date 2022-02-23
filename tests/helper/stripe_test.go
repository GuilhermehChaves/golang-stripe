package helper

import (
	"apiexample/src/helper"
	"testing"
)

func TestCreateStripeRequest(t *testing.T) {
	_, err := helper.CreateStripeRequest(
		"https://stripe.com",
		"https://example.com/success",
		"https://example.com/cancel",
		200.99,
	)

	if err != nil {
		t.Errorf("An error ocurred %s", err)
	}
}
