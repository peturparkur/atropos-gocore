package utils

import (
	"testing"

	"go.uber.org/zap"
)

func TestSimpleGetReq(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()

	apiClient := APIClient{
		BaseURL: "http://postman-echo.com",
		Token:   "",
		Logger:  l,
	}

	resp := map[string]interface{}{}
	if err := apiClient.Get("/get", &resp); err != nil {
		t.Errorf("Error making GET request: %v", err)
	}

	if resp == nil {
		t.Errorf("Response is nil")
	}

	if resp["args"] == nil {
		t.Errorf("Response does not contain args")
	}
	if resp["headers"] == nil {
		t.Errorf("Response does not contain headers")
	}
	if resp["url"] == nil {
		t.Errorf("Response does not contain url")
	}
}

func TestAccidentalNonPointerResp(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()

	apiClient := APIClient{
		BaseURL: "http://postman-echo.com",
		Token:   "",
		Logger:  l,
	}

	resp := interface{}(nil)
	err := apiClient.Get("/get", resp)

	if err == nil {
		t.Errorf("Expected error when passing non-pointer response")
	}

	if devErr, ok := err.(*DeveloperError); ok {
		if devErr.Message != "response provided must be a pointer" {
			t.Errorf("Expected error message to be 'response provided must be a pointer', got %s", devErr.Message)
		}
	} else {
		t.Errorf("Expected error to be of type DeveloperError")
	}
}
