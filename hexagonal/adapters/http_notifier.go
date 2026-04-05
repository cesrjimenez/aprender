package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPNotifier struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPNotifier(baseURL string) *HTTPNotifier {
	return &HTTPNotifier{BaseURL: baseURL, Client: &http.Client{}}
}

func (n *HTTPNotifier) SendConfirmation(email, message string) error {
	body, _ := json.Marshal(map[string]string{
		"to":      email,
		"message": message,
	})
	resp, err := n.Client.Post(n.BaseURL+"/send", "application/json", bytes.NewReader(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notification failed: %w", err)
	}
	return nil
}
