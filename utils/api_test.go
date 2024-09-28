package utils

import (
	"testing"
)

func TestSimpleGetReq(t *testing.T) {
	apiClient := NewAPIClient("http://postman-echo.com", "")

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
	apiClient := NewAPIClient("http://postman-echo.com", "")

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
