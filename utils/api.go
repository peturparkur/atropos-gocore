package gocore

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// MakeGetRequest makes a GET request to the specified endpoint. If token is not "" it will be added to the request as a Bearer token.
func MakeGetRequest[responseType any](l *zap.SugaredLogger, apiBaseURL, endpoint, token string, res *responseType) error {
	req, err := http.NewRequest("GET", apiBaseURL+endpoint, nil)
	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
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

	if err := json.Unmarshal(body, &res); err != nil {
		l.Errorw("Failed to unmarshal response", "error", err, "body", string(body))
		return err
	}

	return nil
}

//
// func makeAPIRequest[respType any](_ *zap.SugaredLogger, token, apiURL, endpoint, kind string, res *respType, msg interface{}) error {
// 	var req *http.Request
// 	var err error
//
// 	if kind == "POST" || kind == "PUT" {
// 		jsonData, err := json.Marshal(msg)
// 		if err != nil {
// 			return err
// 		}
// 		req, err = http.NewRequest(kind, apiURL+endpoint, bytes.NewBuffer(jsonData))
// 	} else {
// 		req, err = http.NewRequest(kind, apiURL+endpoint, nil)
// 	}
//
// 	if err != nil {
// 		return err
// 	}
//
// 	if token != "" {
// 		req.Header.Set("Authorization", "Bearer "+token)
// 	}
// 	if kind == "POST" || kind == "PUT" {
// 		req.Header.Set("Content-Type", "application/json")
// 	}
//
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
//
// 	// Unmarshal the response into a struct
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
