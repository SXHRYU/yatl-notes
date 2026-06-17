package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		User     string
		Password string
		DB       string
		Port     int
	}
	Http struct {
		Host string
		Port int
	}
}

func ReadConfig() (*Config, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	v.SetConfigName(".env")
	v.SetConfigType("env")

	// double check if you're launching from root of the project
	v.AddConfigPath(".")
	// retard alert
	v.AddConfigPath("../../")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
