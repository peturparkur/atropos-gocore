// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
ConsumeWebhookCallback consumes a webhook callback and calls the callback function for each event
The typical usage is as follows:

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		err := ConsumeWebhookCallback(c.Request.Body, func(webhook WebhookCallback) error {
			// Do something with the webhook
			return nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
*/
func ConsumeWebhookCallback(body io.ReadCloser, callback func(webhook WebhookCallback) error) error {
	defer body.Close()
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	stringedEvents := [][]string{}
	if err := json.Unmarshal(bytes, &stringedEvents); err != nil {
		return err
	}

	for _, le := range stringedEvents {
		for _, e := range le {
			event := WebhookCallback{}
			if err := json.Unmarshal([]byte(e), &event); err != nil {
				return err
			}
			if err := callback(event); err != nil {
				return err
			}
		}
	}

	return nil
}

// ConsumeWebhookCallbackWithGin is a helper function to consume a webhook callback with gin
func GinHandler(callback func(webhook WebhookCallback) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := ConsumeWebhookCallback(c.Request.Body, callback)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
