// This package contains the types that are used in the Vikunja API
package vikunja

// VikunjaOwner is a struct that represents an owner in Vikunja
type Owner struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

// VikunjaProject is a struct that represents a project in Vikunja
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

// VikunjaWebhook represents a webhook in Vikunja
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
