package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New(file string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(file)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return &Config{
		v: v,
	}, nil
}

func (c *Config) Viper() *viper.Viper {
	return c.v
}
