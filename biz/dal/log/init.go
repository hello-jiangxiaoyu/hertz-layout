package log

import (
	"github.com/rs/zerolog"
)

var (
	Zero   *zerolog.Logger
	Access *zerolog.Logger
)

func Init() error {
	var err error
	if Zero, err = NewZeroErrorLog(); err != nil {
		return err
	}
	if Access, err = NewZeroAccessLog(); err != nil {
		return err
	}

	return nil
}
