package utils

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// LogFuncErr handle function err
func LogFuncErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		hlog.Error("### Defer err: ", err)
	}
}

// DtoFilter 数组dto
func DtoFilter[S any, T any](s []S, f func(S) T) []T {
	var l []T
	for _, i := range s {
		l = append(l, f(i))
	}
	return l
}
