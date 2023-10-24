package redis

import (
	"context"
	"hertz/demo/pkg/conf"

	"github.com/redis/go-redis/v9"
)

var Rds *redis.Client

func Init() error {
	Rds = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := Rds.Ping(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}
