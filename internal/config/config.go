package config

import (
	"fmt"
	"os"

	"github.com/zalando/go-keyring"
)

type Config struct {
	GitToken string
}

const (
	service = "cobraclip"
	user    = "github-pat"
)

func LoadConfig() (*Config, error) {
	token, err := keyring.Get(service, user)
	if err == nil {
		return &Config{GitToken: token}, nil
	}
	envtoken := os.Getenv("COBRACLIP_GIT_TOKEN")
	//fmt.Printf("DEBUG: Loading config: COBRACLIP_GIT_TOKEN=%s\n", token)

	if envtoken == "" {
		return nil, fmt.Errorf("no token found in COBRACLIP_GIT_TOKEN; please run 'cobraclip login'")
	}

	cfg := &Config{
		GitToken: envtoken,
	}
	return cfg, nil
}

func SaveToken(token string) error {
	return keyring.Set(service, user, token)
}

func HasToken() bool {
	_, err := keyring.Get(service, user)
	return err == nil
}

func SendToken() (string, error) {
	return keyring.Get(service, user)
}

func RemoveToken(token string) error {
	return keyring.Delete(service, user)
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
