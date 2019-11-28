package models

// Experience struct
type Experience struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Lokasi      string `json:"lokasi"`
}

type ExperienceResponse struct {
	Data    []Experience `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}
