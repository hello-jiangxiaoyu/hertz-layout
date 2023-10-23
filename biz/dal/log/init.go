package log

import (
	"github.com/rs/zerolog"
	"time"
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

	Info().Msg("new zero log ok")
	return nil
}

func Info() *zerolog.Event {
	return Zero.Info().Str("ts", time.Now().Format("2006-01-02T15:04:05.000"))
}
