// Package conf 系统相关配置
package conf

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
	"hertz/demo/pkg/utils"
	"os"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz Hertz `yaml:"hertz"`
	MySQL MySQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DB       int    `yaml:"db"`
}

type Hertz struct {
	Address         string `yaml:"address"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
	LogMaxSize      int    `yaml:"log_max_size"`
	LogMaxBackups   int    `yaml:"log_max_backups"`
	LogMaxAge       int    `yaml:"log_max_age"`
	LogFullStack    bool   `yaml:"log_full_stack"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf) // 只读取一次
	return conf
}
func initConf() {
	confFileRelPath := GetConfigPath()
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		hlog.Error("read config file error - %v", err)
		panic(err)
	}

	conf = new(Config)
	if err = yaml.Unmarshal(content, conf); err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}

	hlog.Info("conf init ok: ", utils.StructToString(conf))
}
