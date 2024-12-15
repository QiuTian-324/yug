package config

import (
	"fmt"
	"yug_server/global"

	"github.com/spf13/viper"
)

func LoadFileStorageConfig() (*global.FileStorageConfig, error) {
	var config global.FileStorageConfig
	err := viper.Sub("file_storage").Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("无法解析文件存储配置: %w", err)
	}

	return &config, nil
}
