package models

const (
	HomeAttachment = "localhost:9090/file?path="

	RestoranTypeAttachment   = 1
	KulinerTypeAttachment    = 2
	ExperienceTypeAttachment = 3
)

type Attachment struct {
	AttachmentID int64  `json:"attachment_id"`
	Type         int    `json:"type"`
	Content      string `json:"content"`
}

// type ContenType struct {
// 	ImageUrl []string `json:"image_url"`
// }
