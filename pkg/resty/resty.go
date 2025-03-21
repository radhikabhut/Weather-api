package resty

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

// NewRestClient initializes a new Resty client with optional base URL
func NewRestClient() {
	client = resty.New()
}

// Get sends a GET request
func Get(endpoint string, queryParams map[string]string, headers map[string]string) (*resty.Response, error) {
	req := client.R()
	setHeadersAndParams(req, queryParams, headers)

	resp, err := req.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	return resp, nil
}

// Post sends a POST request with a JSON payload
func Post(endpoint string, body interface{}, headers map[string]string) (*resty.Response, error) {
	req := client.R()
	setHeaders(req, headers)

	resp, err := req.SetBody(body).Post(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %v", err)
	}
	return resp, nil
}

// Put sends a PUT request with a JSON payload
func Put(endpoint string, body interface{}, headers map[string]string) (*resty.Response, error) {
	req := client.R()
	setHeaders(req, headers)

	resp, err := req.SetBody(body).Put(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make PUT request: %v", err)
	}
	return resp, nil
}

// Delete sends a DELETE request
func Delete(endpoint string, headers map[string]string) (*resty.Response, error) {
	req := client.R()
	setHeaders(req, headers)

	resp, err := req.Delete(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to make DELETE request: %v", err)
	}
	return resp, nil
}

// Helper function to set headers and query params
func setHeadersAndParams(req *resty.Request, queryParams map[string]string, headers map[string]string) {
	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}
}

// Helper function to set headers
func setHeaders(req *resty.Request, headers map[string]string) {
	if headers != nil {
		req.SetHeaders(headers)
	}
}
