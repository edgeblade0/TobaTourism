package models

const (
	MessageSuccess  = "Success"
	MessageFailed   = "Failed"
	MessageNotFound = "Not Found"

	StatusSuccess  = "200"
	StatusFailed   = "500"
	StatusNotFound = "404"
)

type DatabaseConfig struct {
	Conn string
}

type Config struct {
	Db DatabaseConfig `mapstructure:"database"`
}

type Responses struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
