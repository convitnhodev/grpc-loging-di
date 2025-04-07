package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	grpc "github.com/convitnhodev/common/grpc"
	logger "github.com/convitnhodev/common/logging"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var defaultConfig []byte

type Config struct {
	Name         string         `mapstructure:"name" yaml:"name"`
	GrpcConfig   *grpc.Config   `mapstructure:"grpc" yaml:"grpc"`
	Loggerconfig *logger.Config `mapstructure:"logger" yaml:"logger"`
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		fmt.Println("Failed to read viper config", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("Failed to unmarshal config", err)
	}

	fmt.Println("Config loaded", cfg)
	return cfg
}
