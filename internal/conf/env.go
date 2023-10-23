package conf

import "os"

const (
	DefaultCookieKey  = "hertz-layout"
	DefaultConfigPath = "internal/conf/conf.online.yaml"
)

func getEnv(env, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetConfigPath() string {
	return getEnv("CONFIG_PATH", DefaultConfigPath)
}
