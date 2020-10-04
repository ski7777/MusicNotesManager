package config

type Config struct {
	Database struct {
		Driver string `json:"driver"`
		DSN    string `json:"dsn"`
	} `json:"database"`
	FileStore string `json:"filestore"`
	WebAddr   string `json:"web"`
}
