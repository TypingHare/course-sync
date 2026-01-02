package config

type Config struct {
	// Student struct{} `mapstructure:"student"`
	// Master struct{} `mapstructure:"master"`
}

func GetDefault() Config {
	var config Config

	return config
}
