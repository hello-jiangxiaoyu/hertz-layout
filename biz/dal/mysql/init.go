package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz/demo/internal/conf"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {
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
