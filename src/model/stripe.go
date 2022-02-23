package model

// type Provider interface {
// 	GetUri(request *CheckoutRequest) (*CheckoutResponse, error)
// }

type CheckoutRequest struct {
	Amount     float32 `json:"amount"`
	SuccessUrl string  `json:"success_url"`
	CancelUrl  string  `json:"cancel_url"`
}

type CheckoutResponse struct {
	Url string `json:"url"`
}
