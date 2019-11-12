package models

// type Pariwisata struct {
// 	ID      int64  `json:"id"`
// 	Title   string `json:"title" validate:"required"`
// 	Content string `json:"content" validate:"required"`
// }

type Pariwisata struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Lokasi string `json:"lokasi"`
}

type Respone struct {
	Data    []Pariwisata `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}
