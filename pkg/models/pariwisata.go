package models

// type Pariwisata struct {
// 	ID      int64  `json:"id"`
// 	Title   string `json:"title" validate:"required"`
// 	Content string `json:"content" validate:"required"`
// }

type Pariwisata struct {
	ID          int64  `json:"id"`
	Nama        string `json:"name"`
	Lokasi      string `json:"lokasi"`
	Description string `json:"desc"`
	Contact     string `json:"contact"`
}

type PariwisataResponse struct {
	Data    []Pariwisata `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}
