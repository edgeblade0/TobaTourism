package models

const (
	PathFileRestoran = "file/restoran/"
	PathFileKuliner  = "file/kuliner/"
)

type Restoran struct {
	RestoID      int64      `json:"restaurantId,omitempty"`
	Name         string     `json:"restaurantName,omitempty"`
	Contact      string     `json:"restaurantContact,omitempty"`
	Location     string     `json:"restaurantLocation,omitempty"`
	ListCulinary []Culinary `json:"culinaryList,omitempty"`
	AttachmentID int64      `json:"attachmentID,omitempty"`
	Attachment   []string   `json:"restaurantImage,omitempty"`
}

type Culinary struct {
	CulinaryID   int64    `json:"culinaryId,omitempty"`
	RestoID      int64    `json:"restaurantId,omitempty"`
	Name         string   `json:"culinaryName,omitempty"`
	Price        int64    `json:"culinaryPrice,omitempty"`
	Desc         string   `json:"culinaryDesc,omitempty"`
	AttachmentID int64    `json:"attachmentID,omitempty"`
	Attachment   []string `json:"culinaryImage,omitempty"`
}
