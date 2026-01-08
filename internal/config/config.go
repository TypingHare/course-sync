package config

type Config struct {
	Student struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"student"`
	// Master struct{} `mapstructure:"master"`
}

func GetDefault() Config {
	var config Config
	return config
}
