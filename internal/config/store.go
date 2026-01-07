package config

import "github.com/spf13/viper"

// Store represents a configuration store using Viper.
type Store struct {
	v *viper.Viper
}
