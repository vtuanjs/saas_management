package viper

import (
	"fmt"

	system "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/core"

	"github.com/spf13/viper"
)

func NewConfig() *system.Config {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading config file, using proceed with environment variables or defaults")
		}
	}

	// Try to read from .env.local if exists
	v.SetConfigName(".env.local")
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading .env.local file, using proceed with environment variables or defaults")
		}
	}

	var config system.Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}

	fmt.Printf("Configuration loaded: %+v\n", config)

	return &config
}
