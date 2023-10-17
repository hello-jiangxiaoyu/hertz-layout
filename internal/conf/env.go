package conf

import "os"

const (
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
