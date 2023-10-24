package utils

import (
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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

// ToUpperCase 字母转大写
func ToUpperCase(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

// GetHashPassword 加密
func GetHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 校验密码
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
