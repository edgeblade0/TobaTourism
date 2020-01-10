package models

const(
	PathFileTransportasi = "file/transportasi"
)

type Transportasi struct {
	ID          int64  `json:"id"`
	Nama        string `json:"nama"`
	Rute        string `json:"rute"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
	Harga       int64  `json:"harga"`
	// Attechment []string `json:"transportasiImage"`
}

type TransportasiResponse struct {
	Data    []Transportasi `json:"data"`
	Message string         `json:"message"`
	Status  string         `json:"status"`
}
