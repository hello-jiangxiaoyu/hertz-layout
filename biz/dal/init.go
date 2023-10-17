package dal

import (
	"github.com/pkg/errors"
	"hertz/demo/biz/dal/mysql"
	"hertz/demo/biz/dal/redis"
)

func Init() error {
	if err := redis.Init(); err != nil {
		return errors.WithMessage(err, "init redis err")
	}
	if err := mysql.Init(); err != nil {
		return errors.WithMessage(err, "init mysql err")
	}
	return nil
}
