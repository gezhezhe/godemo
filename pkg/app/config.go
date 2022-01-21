package app

import "addnewer.com/godemo/pkg/config"

type Config struct {
	Name  string
	Debug bool
	Env   string
}

func NewConfig(c *config.Config) *Config {
	conf := &Config{}
	err := c.Viper().UnmarshalKey("app", conf)
	if err != nil {
		panic("App配置解析错误: " + err.Error())
	}

	return conf
}
