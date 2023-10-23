package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// 日志分片
func timeDivisionWriter(path string) (io.Writer, error) {
	MaxAge := 7
	hook, err := rotate.New(
		path+"_%Y-%m-%d.log",
		rotate.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(MaxAge))),
		rotate.WithRotationTime(time.Minute),
		rotate.WithLinkName(path), // 最新日志软链接
	)

	if err != nil {
		return nil, err
	}
	return hook, nil
}
