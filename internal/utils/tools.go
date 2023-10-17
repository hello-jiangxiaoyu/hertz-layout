package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// DeferErr handle defer function err
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		fmt.Println("### Defer err: ", err)
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

// GetPanicStackInfo 只打印本项目相关的堆栈信息
func GetPanicStackInfo(msg string, err any, skip int, fullStack bool) string {
	pwd, _ := os.Getwd()
	pwd = strings.ReplaceAll(pwd, `\`, "/") // handle windows path
	res := fmt.Sprintf("\n[Recovery] panic recovered: %s\n[Error] %v", msg, err)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		if fullStack || pwd == "" || strings.Contains(file, pwd) { // only about quick auth source files
			res += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
		}
	}
	return res + "\n"
}
