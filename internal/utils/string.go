package utils

import (
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"net"
)

// IsIpAddress 判断是否为ip
func IsIpAddress(host string) bool {
	hostWithoutPort, _, err := net.SplitHostPort(host)
	if err != nil {
		hostWithoutPort = host
	}

	ip := net.ParseIP(hostWithoutPort)
	return ip != nil
}

// StructToString 将结构体转为string，方便打印
func StructToString(obj any) string {
	res, err := make([]byte, 0), errors.New("")
	if res, err = sonic.Marshal(&obj); err != nil {
		return ""
	}
	return string(res)
}

// DefaultValIfEmpty 为空就返回默认值
func DefaultValIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
