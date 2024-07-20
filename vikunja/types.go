// Package vikunja contains the types that are used in the Vikunja API
package vikunja

// Owner is a struct that represents an owner in Vikunja
type Owner struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

// Doer is a struct that represents a doer in Vikunja
type Doer Owner

// Project is a struct that represents a project in Vikunja
type Project struct {
	ID                    int         `json:"id"`
	Title                 string      `json:"title"`
	HexColor              string      `json:"hex_color"`
	BackgroundBlurHash    string      `json:"background_blur_hash"`
	Position              float64     `json:"position"`
	Created               string      `json:"created"`
	Updated               string      `json:"updated"`
	Description           string      `json:"description"`
	DefaultBucketID       int         `json:"default_bucket_id"`
	DoneBucketID          int         `json:"done_bucket_id"`
	Identifier            string      `json:"identifier"`
	IsArchived            bool        `json:"is_archived"`
	IsFavorite            bool        `json:"is_favorite"`
	Owner                 Owner       `json:"owner"`
	ParentProjectID       int         `json:"parent_project_id"`
	BackgroundInformation interface{} `json:"background_information"`
}

// Task represents a task in Vikunja
type Task struct {
	ID                     int         `json:"id"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	Done                   bool        `json:"done"`
	DoneAt                 string      `json:"done_at"`
	DueDate                string      `json:"due_date"`
	Reminders              interface{} `json:"reminders"`
	ProjectID              int         `json:"project_id"`
	RepeatAfter            int         `json:"repeat_after"`
	RepeatMode             int         `json:"repeat_mode"`
	Priority               int         `json:"priority"`
	StartDate              string      `json:"start_date"`
	EndDate                string      `json:"end_date"`
	Assignees              interface{} `json:"assignees"`
	Labels                 interface{} `json:"labels"`
	HexColor               string      `json:"hex_color"`
	PercentDone            int         `json:"percent_done"`
	Identifier             string      `json:"identifier"`
	Index                  int         `json:"index"`
	RelatedTasks           interface{} `json:"related_tasks"`
	Attachments            interface{} `json:"attachments"`
	CoverImageAttachmentID int         `json:"cover_image_attachment_id"`
	IsFavorite             bool        `json:"is_favorite"`
	Created                string      `json:"created"`
	Updated                string      `json:"updated"`
	BucketID               int         `json:"bucket_id"`
	Position               float64     `json:"position"`
	KanbanPosition         float64     `json:"kanban_position"`
	CreatedBy              Owner       `json:"created_by"`
}

// Webhook represents a webhook in Vikunja
type Webhook struct {
	ID        int      `json:"id"`
	TargetURL string   `json:"target_url"`
	Events    []string `json:"events"`
	ProjectID int      `json:"project_id"`
	Secret    *string  `json:"secret,omitempty"`
	Created   *string  `json:"created,omitempty"`
	Updated   *string  `json:"updated,omitempty"`
	CreatedBy *Owner   `json:"created_by,omitempty"`
}

// WebhookCallbackData represents the data that is sent to a webhook callback
type WebhookCallbackData struct {
	Doer Doer `json:"doer"`
	Task Task `json:"task"`
}

// WebhookCallback represents a webhook callback
type WebhookCallback struct {
	EventName string              `json:"event_name"`
	Time      string              `json:"time"`
	Data      WebhookCallbackData `json:"data"`
}
