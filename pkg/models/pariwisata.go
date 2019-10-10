package models

// type Pariwisata struct {
// 	ID      int64  `json:"id"`
// 	Title   string `json:"title" validate:"required"`
// 	Content string `json:"content" validate:"required"`
// }

type Pariwisata struct {
	ID     int64
	Name   string
	Lokasi string
}
