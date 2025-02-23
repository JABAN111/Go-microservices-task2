package config

type Config struct {
	BindPort string `yaml:"port" env:"PETNAME_GRPC_PORT" default:"8080"`
	BindHost string `yaml:"host" env:"PETNAME_GRPC_HOST" default:"0.0.0.0"`
}

func NewConfig() *Config {
	return &Config{
		BindPort: "",
		BindHost: "0.0.0.0", // NOTE: костылек, в задании не задается хост
	}
}

func DefaultConfig() *Config {
	return &Config{
		BindPort: "8080",
		BindHost: "0.0.0.0",
	}
}
