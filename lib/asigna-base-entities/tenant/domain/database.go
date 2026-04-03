package domain

type DatabaseConfig struct {
	TenantId string
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
	SslMode  string
	Settings map[string]any
}
