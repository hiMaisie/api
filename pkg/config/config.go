package config

type Config struct {
	HTTP *HTTPConfig
}

type HTTPConfig struct {
	HostName string
	Port     uint
}
