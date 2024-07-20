// Package utils contains utility functions that are used throughout the application.
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"go.uber.org/zap"
)

// APIError is an error type that is returned when an API request fails.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return "API request failed with status code " + strconv.Itoa(e.StatusCode) + ": " + e.Message
}

// MakeAPIRequest is a generic function to make an API request. It supports GET, POST, PUT, and DELETE requests.
func MakeAPIRequest(l *zap.SugaredLogger, kind, apiBaseURL, endpoint, token string, request, response interface{}) error {
	// If response is not nil, check its a pointer (easy dev mistake to make).
	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		return &DeveloperError{"response provided must be a pointer"}
	}

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

	var req *http.Request

	if request != nil {
		buf := bytes.NewBuffer(jsonData)
		req, err = http.NewRequest(kind, apiBaseURL+endpoint, buf)
	} else {
		req, err = http.NewRequest(kind, apiBaseURL+endpoint, nil)
	}

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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{StatusCode: resp.StatusCode, Message: string(body)}
	}

	return nil
}

// MakeDeleteRequest is a helper function to make a DELETE request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakeDeleteRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, response interface{}) error {
	return MakeAPIRequest(l, "DELETE", apiBaseURL, endpoint, token, nil, response)
}

// MakeGetRequest is a helper function to make a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakeGetRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, response interface{}) error {
	return MakeAPIRequest(l, "GET", apiBaseURL, endpoint, token, nil, response)
}

// MakePostRequest is a helper function to make a POST request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakePostRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, request, response interface{}) error {
	return MakeAPIRequest(l, "POST", apiBaseURL, endpoint, token, request, response)
}

// MakePutRequest is a helper function to make a PUT request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakePutRequest(l *zap.SugaredLogger, apiBaseURL, endpoint, token string, request, response interface{}) error {
	return MakeAPIRequest(l, "PUT", apiBaseURL, endpoint, token, request, response)
}
