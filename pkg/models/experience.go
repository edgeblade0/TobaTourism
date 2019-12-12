package models

const (
	PathFileExperience = "file/experience/"
)

// Experience struct
type Experience struct {
	ID           int64    `json:"id,omitempty"`
	Description  string   `json:"description,omitempty"`
	Location     string   `json:"location,omitempty"`
	AttachmentID int64    `json:"attachment_id,omitempty"`
	Attachments  []string `json:"attachments,omitempty"`
}

type ExperienceResponse struct {
	Data    []Experience `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
	Status  string       `json:"status,omitempty"`
}
