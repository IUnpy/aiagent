package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config 应用配置结构
type Config struct {
	APIKey string `json:"api_key"`
	// 可以添加其他配置项
	Theme struct {
		Dark bool `json:"dark"`
	} `json:"theme"`
	Window struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"window"`
}

var defaultConfig = Config{
	APIKey: "sk-sbscazwmhwgqdshepehxtrfwwoavvfmxfskspscjuwwzjowr",
	Theme: struct {
		Dark bool `json:"dark"`
	}{
		Dark: true,
	},
	Window: struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{
		Width:  500,
		Height: 400,
	},
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return &defaultConfig, err
	}

	appConfigDir := filepath.Join(configDir, "aiagent")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return &defaultConfig, err
	}

	configFile := filepath.Join(appConfigDir, "config.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果配置文件不存在，创建默认配置
			return SaveConfig(&defaultConfig)
		}
		return &defaultConfig, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &defaultConfig, err
	}

	return &config, nil
}

// SaveConfig 保存配置
func SaveConfig(config *Config) (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return config, err
	}

	appConfigDir := filepath.Join(configDir, "aiagent")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return config, err
	}

	configFile := filepath.Join(appConfigDir, "config.json")
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return config, err
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return config, err
	}

	return config, nil
}
