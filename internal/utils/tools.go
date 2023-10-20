package utils

import (
	"fmt"
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
