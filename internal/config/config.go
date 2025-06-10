package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	gittoken string
}

func loadconfig() (*Config, error) {
	homedir, err := os.UserHomeDir()

	if err != nil {
		return nil, fmt.Errorf("failed to get homedir", err)
	}

	configdir := filepath.Join(homedir, ".cobraclip")
	filepath.Join(configdir, "config.yaml")

	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configdir)

	viper.SetDefault("git_token", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
		if err := os.MkdirAll(configdir, 0700); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	viper.SetEnvPrefix("cobraclip")
	viper.BindEnv("git_token")
	viper.AutomaticEnv()

	cfg := &Config{
		gittoken: viper.GetString("git_token"),
	}

	return cfg, nil
}

func saveToken(token string) error {
	viper.Set("git_token", token)

	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return viper.SafeWriteConfig()
		}

		return fmt.Errorf("failed to write into config.yaml", err)
	}

	return nil
}
