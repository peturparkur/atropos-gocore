// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/atropos112/gocore/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
func ConsumeWebhookCallback(l *zap.SugaredLogger, body io.ReadCloser, callback func(webhook WebhookCallback) error) error {
	defer body.Close()
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	// Figure out if it has body and i need to unpack it or its already unpacked
	obj := map[string]interface{}(nil)
	if err := json.Unmarshal(bytes, &obj); err != nil {
		l.Errorw("Failed to unmarshal webhook event", "event", string(bytes))
		return err
	}

	event := WebhookCallback{}
	if b, ok := obj["body"]; !ok {
		bStr := b.(string)

		if err := json.Unmarshal([]byte(bStr), &event); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal(bytes, &event); err != nil {
			return err
		}
	}

	if err := callback(event); err != nil {
		return err
	}

	return nil
}

// GinHandler is a helper function to consume a webhook callback with gin
func GinHandler(l *zap.SugaredLogger, callback func(webhook WebhookCallback) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := ConsumeWebhookCallback(l, c.Request.Body, callback)
		if err != nil {
			if devErr, ok := err.(*utils.DeveloperError); ok {
				l.Error(devErr)
				c.JSON(http.StatusInternalServerError, gin.H{"error": devErr.Error()})
				return
			}

			l.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
