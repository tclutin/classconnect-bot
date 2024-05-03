package main

import (
	"bytes"
	"classconnect-bot/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	client  http.Client
	baseURL string
	jwt     string
}

func New(baseURL string) *HTTPClient {
	return &HTTPClient{
		client:  http.Client{},
		baseURL: baseURL,
	}
}

func (h *HTTPClient) LogIn(username string, password string) error {
	url := h.baseURL + "/api/v1/auth/login"
	payload := model.LoginRequest{
		Username: username,
		Password: password,
	}

	jsonPayload, _ := json.Marshal(payload)

	resp, err := h.client.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status code: %d", resp.StatusCode)
	}

	var token model.TokenResponse

	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return err
	}

	h.jwt = token.AccessToken

	return nil
}

func (h *HTTPClient) Validate(username string, password string) error {

}

func main() {
	err := New("http://127.0.0.1:8080").LogIn("telegram", "telegram")
	if err != nil {
		fmt.Println(err)
	}
}
