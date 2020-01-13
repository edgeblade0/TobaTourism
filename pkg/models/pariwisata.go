package models

const (
	PathFilePariwisata = "file/pariwisata/"
)

type Pariwisata struct {
	ID           int64    `json:"tourismId,omitempty"`
	Nama         string   `json:"tourismName,omitempty"`
	Lokasi       string   `json:"tourismLocation,omitempty"`
	Description  string   `json:"tourismDescription,omitempty"`
	Contact      string   `json:"tourismContact,omitempty"`
	AttachmentID int64    `json:"attachmentID,omitempty"`
	Attachment   []string `json:"tourismImage,omitempty"`
}

type PariwisataResponse struct {
	Data    []Pariwisata `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}

type PariwisataResponseByID struct {
	Data    Pariwisata `json:"data"`
	Message string     `json:"message"`
	Status  string     `json:"status"`
}
