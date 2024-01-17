package delivery

import (
	"bytes"
	"fmt"
	"net/http"
)

type Delivery interface {
	Send(data []byte) error
}

type HttpsDelivery struct {
	Endpoint string
}

func NewHttpsDelivery(endpoint string) *HttpsDelivery {
	return &HttpsDelivery{
		Endpoint: endpoint,
	}
}

func (h *HttpsDelivery) Send(data []byte) error {

	resp, err := http.Post(h.Endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	return nil
}
