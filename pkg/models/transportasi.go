package models

const (
	PathFileTransportasi = "file/transportasi/"
)

type Transportasi struct {
	ID           int64    `json:"transportationId"`
	Nama         string   `json:"transportationName"`
	Rute         string   `json:"transportationRoute"`
	Description  string   `json:"transportationDescription"`
	Contact      string   `json:"transportationContact"`
	Harga        int64    `json:"transportationPrice"`
	Attachment   []string `json:"transportationImage"`
	AttachmentID int64    `json:"attachmentId, omitempty"`
}

type TransportasiResponse struct {
	Data    []Transportasi `json:"data,omitempty"`
	Message string         `json:"message"`
	Status  string         `json:"status"`
}

type TransportasiDetailResponse struct {
	Data    Transportasi `json:"data,omitempty"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}
