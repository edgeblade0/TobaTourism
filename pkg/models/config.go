package models

type DatabaseConfig struct {
	Conn string
}

type Config struct {
	Db DatabaseConfig `mapstructure:"database"`
}
