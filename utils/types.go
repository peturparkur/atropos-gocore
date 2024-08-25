// Package utils contains utility functions that are used throughout the application.
package utils

import (
	"fmt"
)

// APIClient is a struct that contains the base URL of the API and the token to use for requests.
type APIClient struct {
	BaseURL string
	Token   string
}

// Delete is a helper function to make a DELETE request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Delete(endpoint string, response interface{}) error {
	return MakeDeleteRequest(c.BaseURL, endpoint, c.Token, response)
}

// Get is a helper function to make a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Get(endpoint string, response interface{}) error {
	return MakeGetRequest(c.BaseURL, endpoint, c.Token, response)
}

// Post is a helper function to make a POST request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Post(endpoint string, request, response interface{}) error {
	return MakePostRequest(c.BaseURL, endpoint, c.Token, request, response)
}

// Put is a helper function to make a PUT request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *APIClient) Put(endpoint string, request, response interface{}) error {
	return MakePutRequest(c.BaseURL, endpoint, c.Token, request, response)
}

// NoCredFoundError represents an error when no credentials are found
type NoCredFoundError struct {
	CredentialName string
}

func (e *NoCredFoundError) Error() string {
	return fmt.Sprintf("no credentials found for %s", e.CredentialName)
}

type DeveloperError struct {
	Message string
}

func (e *DeveloperError) Error() string {
	return fmt.Sprintf("developer error: %s", e.Message)
}

type GPTDoesntListenError struct {
	UserMessage string
	SysMessage  string
}

func (e *GPTDoesntListenError) Error() string {
	// Write sys message and user message to log and return
	return fmt.Sprintf("GPT doesn't listen, sys message: %s, user message: %s", e.SysMessage, e.UserMessage)
}
