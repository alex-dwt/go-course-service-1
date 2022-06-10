package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type YamlConfig struct {
	Author    string         `validate:"oneof=alex vasya"`
	DBCfg     DBConfig       `mapstructure:"db"`
	RedisCfg  RedisConfig    `mapstructure:"redis"`
	MtConfigs []MtConfigItem `validate:"dive"`
}

type DBConfig struct {
	Port     int
	Host     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Port int `validate:"gt=6665"`
	Host string
}

type MtConfigItem struct {
	Type     string
	Leverage int `validate:"gt=5,lte=100"`
}

func NewYamlConfig(configFilePath string) (YamlConfig, error) {
	v := viper.New()
	v.SetConfigFile(configFilePath)

	if err := v.ReadInConfig(); err != nil {
		return YamlConfig{}, fmt.Errorf("viper read: %w", err)
	}

	var c YamlConfig
	if err := v.Unmarshal(&c); err != nil {
		return YamlConfig{}, fmt.Errorf("viper unmarshal: %w", err)
	}

	if err := validateYamlConfig(c); err != nil {
		return YamlConfig{}, fmt.Errorf("validation: %w", err)
	}

	return c, nil
}
