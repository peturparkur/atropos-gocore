// Package utils contains utility functions that are used throughout the application.
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// APIClient is a struct that contains the base URL of the API and the token to use for requests.
type APIClient struct {
	BaseURL string
	Token   string
	Logger  *zap.SugaredLogger
}

// NewAPIClient creates a new APIClient with the specified baseURL and token.
func NewAPIClient(baseURL, token string, logger *zap.SugaredLogger) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Token:   token,
		Logger:  logger,
	}
}

// NewAPIClient creates a new APIClient with the specified baseURL and token.
func (c *APIClient) Delete(endpoint string, response interface{}) error {
	return MakeDeleteRequest(c.Logger, c.BaseURL, endpoint, c.Token, response)
}

// Get is a helper function to make a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Get(endpoint string, response interface{}) error {
	return MakeGetRequest(c.Logger, c.BaseURL, endpoint, c.Token, response)
}

// Post is a helper function to make a POST request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Post(endpoint string, request, response interface{}) error {
	return MakePostRequest(c.Logger, c.BaseURL, endpoint, c.Token, request, response)
}

// Put is a helper function to make a PUT request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Put(endpoint string, request, response interface{}) error {
	return MakePutRequest(c.Logger, c.BaseURL, endpoint, c.Token, request, response)
}

// MakeAPIRequest is a generic function to make an API request. It supports GET, POST, PUT, and DELETE requests.
func MakeApiRequest(l *zap.SugaredLogger, kind, apiBaseURL, endpoint, token string, request, response interface{}) error {
	var jsonData []byte
	var err error

	if request != nil {
		if kind == "GET" || kind == "DELETE" {
			l.Fatalw("GET and DELETE requests do not support request bodies", "kind", kind)
		}

		jsonData, err = json.Marshal(request)
		if err != nil {
			return err
		}
	}

	var buf *bytes.Buffer
	if request != nil {
		buf = bytes.NewBuffer(jsonData)
	} else {
		buf = nil
	}

	req, err := http.NewRequest(kind, apiBaseURL+endpoint, buf)
	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	if request != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		l.Errorw("Failed to unmarshal response", "error", err, "body", string(body))
		return err
	}

	return nil
}

// MakeDeleteRequest is a helper function to make a DELETE request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakeDeleteRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, response interface{}) error {
	return MakeApiRequest(l, "DELETE", apiBaseURL, endpoint, token, nil, response)
}

// MakeGetRequest is a helper function to make a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakeGetRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, response interface{}) error {
	return MakeApiRequest(l, "GET", apiBaseURL, endpoint, token, nil, response)
}

// MakePostRequest is a helper function to make a POST request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakePostRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, request, response interface{}) error {
	return MakeApiRequest(l, "POST", apiBaseURL, endpoint, token, request, response)
}

// MakePutRequest is a helper function to make a PUT request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakePutRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, request, response interface{}) error {
	return MakeApiRequest(l, "PUT", apiBaseURL, endpoint, token, request, response)
}
