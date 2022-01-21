package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDialector(conf *Config) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: conf.Dns,
	})
}
