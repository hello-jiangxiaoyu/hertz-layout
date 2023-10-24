package dal

import "hertz/demo/biz/dal/log"

func Init() error {
	if err := log.Init(); err != nil {
		return err
	}
	//if err := redis.Init(); err != nil {
	//	return errors.WithMessage(err, "init redis err")
	//}
	//if err := mysql.Init(); err != nil {
	//	return errors.WithMessage(err, "init mysql err")
	//}
	return nil
}

func InitAdmin() error {
	//if err := redis.Init(); err != nil {
	//	return errors.WithMessage(err, "init redis err")
	//}
	//if err := mysql.Init(); err != nil {
	//	return errors.WithMessage(err, "init mysql err")
	//}
	return nil
}
