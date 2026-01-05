// Package api provides the http client and wrapper for the AnythingLLM API
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"anything-tui/config"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) handleResponse(resp *http.Response, result any) error {
	if resp.StatusCode >= 400 {
		return c.parseError(resp)
	}

	if result == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("could not unmarshal response body %w", err)
	}

	return nil
}

func (c *Client) parseError(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)

	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.ErrorMessage() != "unknown error" {
		apiErr.StatusCode = resp.StatusCode
		return fmt.Errorf("API Error (%d) %s", apiErr.StatusCode, apiErr.ErrorMessage())
	}

	return fmt.Errorf("API Error (%d) %s", resp.StatusCode, string(body))
}

func (c *Client) request(method string, path string, body any) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	url := c.baseURL + "/api/v1" + path

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed generating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	return c.httpClient.Do(req)
}

// TODO: Implement func (c *Client) streamRequest(path string, body any) (*http.Response, error) {}

func (c *Client) get(path string, result any) error {
	resp, err := c.request("GET", path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.handleResponse(resp, result)
}

func (c *Client) post(path string, body any, result any) error {
	resp, err := c.request("POST", path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.handleResponse(resp, result)
}

func (c *Client) delete(path string) error {
	resp, err := c.request("DELETE", path, nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return c.parseError(resp)
	}
	return nil
}
