package config

import (
	"sync"
)

var (
	globalConfig *Config
	configMutex  sync.Mutex
)

type Config struct {
	UserName string
	PassWord string
	Age      int
	Name     string
}

func SetConfig(config Config) {
	globalConfig = &config
}

func GetConfig() *Config {
	if globalConfig != nil {
		return globalConfig
	}
	configMutex.Lock()
	defer configMutex.Unlock()

	//第二次检查，防止并发时，多次初始化
	if globalConfig != nil {
		return globalConfig
	}

	return globalConfig
}
