package config

import (
	"fmt"
	"os"
)

type Config struct {
	GitToken string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("COBRACLIP_GIT_TOKEN")
	//fmt.Printf("DEBUG: Loading config: COBRACLIP_GIT_TOKEN=%s\n", token)

	if token == "" {
		return nil, fmt.Errorf("no token found in COBRACLIP_GIT_TOKEN; please run 'cobraclip login'")
	}

	cfg := &Config{
		GitToken: token,
	}
	return cfg, nil
}

// func SaveToken(token string) error {

// 	viper.Set("git_token", token)

// 	// if err := viper.WriteConfig(); err != nil {
// 	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 	// 		return viper.SafeWriteConfig()
// 	// 	}

// 	// 	return fmt.Errorf("failed to write into config.yaml")
// 	// }

// 	return nil
// }
