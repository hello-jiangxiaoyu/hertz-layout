package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	AccessPath = "data/access"
	ErrorPath  = "data/error"
)

func NewZeroAccessLog() (*zerolog.Logger, error) {
	writer, err := timeDivisionWriter(AccessPath)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(writer).With().Logger()
	return &logger, nil
}

func NewZeroErrorLog() (*zerolog.Logger, error) {
	writer, err := timeDivisionWriter(ErrorPath)
	if err != nil {
		return nil, err
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:     writer,
		NoColor: true,
		FormatTimestamp: func(i interface{}) string {
			return ""
		},
	}
	logger := log.Output(consoleWriter).With().Caller().Logger()
	return &logger, nil
}
