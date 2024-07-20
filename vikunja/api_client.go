// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"strconv"

	"github.com/atropos112/gocore/utils"
	"go.uber.org/zap"
)

// Client is the interface for the Vikunja API client
type Client utils.APIClient

// GetVikunjaAPIClient returns a new Vikunja API client
func GetVikunjaAPIClient(token, apiURL string) (*Client, error) {
	// Setting up logging
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()
	var err error

	// Get creds
	if token == "" {
		token, err = utils.GetCred(l, "ATRO_VIKUNJA_ATROPOS_API_TOKEN")
		if err != nil {
			return nil, err
		}
	}
	if apiURL == "" {
		apiURL, err = utils.GetCred(l, "ATRO_VIKUNJA_API_URL")
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		BaseURL: apiURL,
		Token:   token,
		Logger:  l,
	}, nil
}

// GetProjects returns a list of projects
func (c *Client) GetProjects() ([]Project, error) {
	apiClient := utils.APIClient(*c)
	projects := []Project{}

	if err := apiClient.Get("/projects", &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// GetProjectWebhooks returns a list of webhooks for a project
func (c *Client) GetProjectWebhooks(projectID int) ([]Webhook, error) {
	apiClient := utils.APIClient(*c)
	webhooks := []Webhook{}
	projectIDstr := strconv.Itoa(projectID)

	if err := apiClient.Get("/projects/"+projectIDstr+"/webhooks", &webhooks); err != nil {
		return nil, err
	}

	return webhooks, nil
}

// CreateProjectWebhook creates a webhook for a project
func (c *Client) CreateProjectWebhook(projectID int, webhook Webhook) (Webhook, error) {
	apiClient := utils.APIClient(*c)
	webhookRes := Webhook{}
	projectIDstr := strconv.Itoa(projectID)

	if err := apiClient.Post("/projects/"+projectIDstr+"/webhooks", webhook, &webhookRes); err != nil {
		return Webhook{}, err
	}

	return webhookRes, nil
}

// UpdateProjectWebhook updates a webhook for a project, only can update events (nothing else)
func (c *Client) UpdateProjectWebhook(projectID int, webhook Webhook) (Webhook, error) {
	apiClient := utils.APIClient(*c)
	projectIDstr := strconv.Itoa(projectID)

	res := Webhook{}
	err := apiClient.Post("/projects/"+projectIDstr+"/webhooks/"+strconv.Itoa(webhook.ID), webhook, &res)

	return res, err
}
