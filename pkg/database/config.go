package database

import (
	"addnewer.com/godemo/pkg/config"
)

type Config struct {
	Dns string
}

func LoadConfig(c *config.Config) (*Config, error) {
	conf := &Config{}
	err := c.Viper().UnmarshalKey("mysql.default", conf)
	if err != nil {
		// panic("数据库配置解析错误: " + err.Error())
		return nil, err
	}

	return conf, nil
}
