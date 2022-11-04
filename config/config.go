package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"crypto_price_tracker/config/support"
	"github.com/spf13/viper"
)

const (
	Development = "development"
	Production  = "production"
	Staging     = "staging"
)

var env string
var configProvider = viper.New()

func init() {
	Init()
}

func Init() {
	// Derive the config directory
	configPath := support.GetCryptoRootDir() + "/config"
	err := SetCryptoConfig(configPath)
	if err != nil {
		panic(err)
	}
	env = GetEnv()
	fmt.Fprintln(os.Stdout, fmt.Sprintf(`{"timestamp":"%s", "env":%s}`,
		time.Now().Format("2006-01-02T15:04:05.000-07:00"), env))
}

func SetCryptoConfig(configPath string) error {
	configProvider.SetConfigName("config")
	configProvider.SetConfigFile(configPath + "/" + "development.json")
	return configProvider.ReadInConfig()
}

// IsDevelopment Returns true if current environment is Development
func IsDevelopment() bool {
	return env == Development
}

// IsProduction Returns true if current environment is Production
func IsProduction() bool {
	return env == Production
}

// IsStaging Returns true if current environment is Staging
func IsStaging() bool {
	return env == Staging
}

func GetEnv() string {
	environment := ENVIRONMENT.Get()
	if strings.Contains(environment, "develop") {
		return Development
	} else if strings.Contains(environment, "staging") {
		return Staging
	}
	return Production
}
