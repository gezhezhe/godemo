package database

import (
	"addnewer.com/godemo/pkg/config"
	"gorm.io/gorm"
)

func NewOrmClient(c *config.Config) (*gorm.DB, error) {
	conf, err := LoadConfig(c)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(NewMysqlDialector(conf))
	if err != nil {
		return nil, err
	}

	return db, err
}
