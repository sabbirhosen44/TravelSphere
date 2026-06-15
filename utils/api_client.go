package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	Client  *http.Client
	Timeout time.Duration
}

func NewAPIClient(timeout time.Duration) *APIClient {
	return &APIClient{
		Client: &http.Client{
			Timeout: timeout,
		},
		Timeout: timeout,
	}
}

// Get performs a GET request to the specified URL with custom headers.
func (c *APIClient) Get(url string, headers map[string]string) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "TravelSphereApp/1.0 (Go-Backend)")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, resp.StatusCode, nil
}
