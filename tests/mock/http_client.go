package mock

import (
	"log"
	"net/http"
)

type HttpClientMock struct{}

func (h *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func (h *HttpClientMock) CloseIdleConnections() {
	log.Println("Closing idle connections")
}
