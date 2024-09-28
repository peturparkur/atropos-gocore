// Package utils contains utility functions that are used throughout the application.
package utils

import (
	"fmt"
	"net/http"
)

// AuthenticatedAPIClient is a struct that contains the base URL of the API and the token to use for requests.
type AuthenticatedAPIClient struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

// NewAPIClient creates a new AuthenticatedAPIClient with the specified base URL and token.
func NewAPIClient(baseURL, token string) AuthenticatedAPIClient {
	return AuthenticatedAPIClient{
		BaseURL: baseURL,
		Token:   token,
		Client:  http.DefaultClient,
	}
}

// Delete is a helper function to make a DELETE request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *AuthenticatedAPIClient) Delete(endpoint string, response interface{}) error {
	return MakeDeleteRequest(c.Client, c.BaseURL, endpoint, c.Token, response)
}

// Get is a helper function to make a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *AuthenticatedAPIClient) Get(endpoint string, response interface{}) error {
	return MakeGetRequest(c.Client, c.BaseURL, endpoint, c.Token, response)
}

// Post is a helper function to make a POST request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *AuthenticatedAPIClient) Post(endpoint string, request, response interface{}) error {
	return MakePostRequest(c.Client, c.BaseURL, endpoint, c.Token, request, response)
}

// Put is a helper function to make a PUT request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func (c *AuthenticatedAPIClient) Put(endpoint string, request, response interface{}) error {
	return MakePutRequest(c.Client, c.BaseURL, endpoint, c.Token, request, response)
}

// NoCredFoundError represents an error when no credentials are found
type NoCredFoundError struct {
	CredentialName string
}

func (e *NoCredFoundError) Error() string {
	return fmt.Sprintf("no credentials found for %s", e.CredentialName)
}

// DeveloperError represents an error that is caused by a developer mistake
type DeveloperError struct {
	Message string
}

func (e *DeveloperError) Error() string {
	return fmt.Sprintf("developer error: %s", e.Message)
}

// GPTDoesntListenError represents an error when GPT doesn't listen
type GPTDoesntListenError struct {
	UserMessage string
	SysMessage  string
}

func (e *GPTDoesntListenError) Error() string {
	// Write sys message and user message to log and return
	return fmt.Sprintf("GPT doesn't listen, sys message: %s, user message: %s", e.SysMessage, e.UserMessage)
}
