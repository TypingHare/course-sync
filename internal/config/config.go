package config

// Config represents the application configuration structure.
type Config struct {
	// Student related configuration.
	Student struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"student"`

	// Master (teachers, professors, TAs, graders, etc.) related configuration.
	Master struct{} `mapstructure:"master"`
}

// GetDefault returns the default configuration.
func GetDefault() Config {
	var config Config
	return config
}
