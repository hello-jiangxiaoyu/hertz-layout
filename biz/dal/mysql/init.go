package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz/demo/pkg/conf"
)

var (
	DB *gorm.DB
)

func Init() error {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
