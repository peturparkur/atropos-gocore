// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"strconv"

	"github.com/atropos112/gocore/utils"
)

// Client is the interface for the Vikunja API client
type Client utils.APIClient

// GetVikunjaAPIClient returns a new Vikunja API client
func GetVikunjaAPIClient(token, apiURL string) (*Client, error) {
	// Setting up logging
	var err error

	// Get creds
	if token == "" {
		token, err = utils.GetCred("ATRO_VIKUNJA_ATROPOS_API_TOKEN")
		if err != nil {
			return nil, err
		}
	}
	if apiURL == "" {
		apiURL, err = utils.GetCred("ATRO_VIKUNJA_API_URL")
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		BaseURL: apiURL,
		Token:   token,
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

// GetProject returns a specific project
func (c *Client) GetProject(projectID int) (Project, error) {
	apiClient := utils.APIClient(*c)
	project := Project{}
	projectIDstr := strconv.Itoa(projectID)

	if err := apiClient.Get("/projects/"+projectIDstr, &project); err != nil {
		return Project{}, err
	}

	return project, nil
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

	if err := apiClient.Put("/projects/"+projectIDstr+"/webhooks", webhook, &webhookRes); err != nil {
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

// GetProjectTasks returns a list of tasks for a project
func (c *Client) GetProjectTasks(projectID int) ([]Task, error) {
	apiClient := utils.APIClient(*c)
	tasks := []Task{}
	projectIDstr := strconv.Itoa(projectID)

	if err := apiClient.Get("/projects/"+projectIDstr+"/tasks", &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateProject updates a project
func (c *Client) UpdateProject(project Project) (Project, error) {
	apiClient := utils.APIClient(*c)
	projectIDstr := strconv.Itoa(project.ID)

	res := Project{}
	err := apiClient.Post("/projects/"+projectIDstr, project, &res)

	return res, err
}

// GetTaskComments returns a list of comments for a task
func (c *Client) GetTaskComments(taskID int) ([]Comment, error) {
	apiClient := utils.APIClient(*c)
	comments := []Comment{}
	taskIDstr := strconv.Itoa(taskID)

	if err := apiClient.Get("/tasks/"+taskIDstr+"/comments", &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetTask returns a task
func (c *Client) GetTask(taskID int) (Task, error) {
	apiClient := utils.APIClient(*c)
	task := Task{}
	taskIDstr := strconv.Itoa(taskID)

	if err := apiClient.Get("/tasks/"+taskIDstr, &task); err != nil {
		return Task{}, err
	}

	return task, nil
}

// UpdateTask updates a task
func (c *Client) UpdateTask(task Task) (Task, error) {
	apiClient := utils.APIClient(*c)
	resp := Task{}
	taskIDstr := strconv.Itoa(task.ID)

	err := apiClient.Post("/tasks/"+taskIDstr, task, &resp)

	return resp, err
}

func (c *Client) GetAllLabels() ([]Label, error) {
	apiClient := utils.APIClient(*c)
	labels := []Label{}

	// per_page is limited up to 50 (default is 50) so need to collect all pages
	pageCount := 1
	for {

		pagelabels := []Label{}
		if err := apiClient.Get("/labels?page="+strconv.Itoa(pageCount), &pagelabels); err != nil {
			return nil, err
		}
		if len(pagelabels) == 0 {
			break
		}
		labels = append(labels, pagelabels...)
		pageCount++
	}

	return labels, nil
}

func (c *Client) AddLabelToTask(taskID, labelID int) (LabelID, error) {
	apiClient := utils.APIClient(*c)
	taskIDstr := strconv.Itoa(taskID)

	req := LabelID{
		ID: labelID,
	}
	resp := LabelID{}
	err := apiClient.Put("/tasks/"+taskIDstr+"/labels", req, &resp)

	return resp, err
}
