package models

const (
	PathFilePariwisata = "file/pariwisata/"
)

type Pariwisata struct {
	ID          int64  `json:"id,omitempty"`
	Nama        string `json:"name,omitempty"`
	Lokasi      string `json:"lokasi,omitempty"`
	Description string `json:"desc,omitempty"`
	Contact     string `json:"contact,omitempty"`
	AttachmentID int64      `json:"attachmentID,omitempty"`
	Attachment   []string   `json:"pariwisataImage,omitempty"`
}

type PariwisataResponse struct {
	Data    []Pariwisata `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}
