package configs

// Config ...
type Config struct {
	LogLevel string `toml:"log_level"`
	Apitoken string `toml:"token"`
}

// New Config
func NewConfig() *Config {
	return &Config{
		LogLevel: "debug",
	}
}
