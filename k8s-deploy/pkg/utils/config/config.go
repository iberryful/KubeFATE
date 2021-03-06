package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const cUserSpecifiedPATH string = "FATECLOUD_CONFIG_PATH"
const cEnvironmentPrefix string = "FATECLOUD"

// DirExists checks if a dir is existed
func DirExists(configPath string) bool {
	fi, err := os.Stat(configPath)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// InitViper initial a viper instance
func InitViper() error {
	// For environment variable
	viper.SetEnvPrefix("FATECLOUD")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// User specified config path, "" by default
	var altPath = os.Getenv(cUserSpecifiedPATH)
	if altPath != "" {
		if !DirExists(altPath) {
			return fmt.Errorf("%s %s does not exist", cUserSpecifiedPATH, altPath)
		}
		viper.AddConfigPath(altPath)
	} else {
		// Append the project dir to the config seraching path
		path, _ := filepath.Abs(".")
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return nil
}

// InitConfig initials the viper and read config in
func InitConfig() error {
	err := InitViper()
	if err != nil {
		return err
	}
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
