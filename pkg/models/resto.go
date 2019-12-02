package models

const (
	PathFileRestoran = "file/restoran/"
	PathFileKuliner  = "file/kuliner/"
)

type Restoran struct {
	RestoID      int64      `json:"resto_id"`
	Contact      string     `json:"contact"`
	Location     string     `json:"location"`
	Name         string     `json:"name"`
	ListCulinary []Culinary `json:"list_culinary"`
	AttachmentID int64      `json:"attachment_id"`
	Attachment   Attachment `json:"attachment"`
}

type Culinary struct {
	CulinaryID   int64      `json:"culinary_id"`
	RestoID      int64      `json:"resto_id"`
	Name         string     `json:"name"`
	Price        int64      `json:"price"`
	Desc         string     `json:"desc"`
	AttachmentID int64      `json:"attachment_id"`
	Attachment   Attachment `json:"attachment"`
}
