package config

import (
	"errors"
	"strconv"
)

type configKey string

const (
	ENVIRONMENT    = configKey("environment")
	BASE_API_URL   = configKey("base.api")
	SENTRY_DSN_KEY = configKey("sentry.dsn")
	SENTRY_ENABLED = configKey("sentry.enable")

	PLOG_LEVEL           = configKey("plog.log_level")
	PLOG_TYPE            = configKey("plog.log_type")
	PLOG_LOCATION        = configKey("plog.log_location")
	CRYPTO_URL           = configKey("crypto_url")
	REDIS_URL            = configKey("redis_url")
	REDIS_PASSWORD       = configKey("redis_password")
	REDIS_EXPIRY_SECONDS = configKey("redis_expiry_seconds")
)

// Gets the value for given key from the config file.
// It panics no configuration value is present
func (c configKey) Get() string {
	val := configProvider.GetString(string(c))
	if val == "" && c != "redis_password" {
		panic(errors.New("Configuration value not found [" + string(c) + "]"))
	}
	return val
}

func (c configKey) GetInt() int {
	val := configProvider.GetString(string(c))
	if val == "" {
		panic(errors.New("Configuration value not found [" + string(c) + "]"))
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		panic(errors.New("Configuration value not int [" + string(c) + "]"))
	}
	return num
}

func (c configKey) GetBool() bool {
	val := configProvider.GetString(string(c))
	if val == "" {
		panic(errors.New("Configuration value not found [" + string(c) + "]"))
	}
	retVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(errors.New("Configuration value not bool [" + string(c) + "]"))
	}
	return retVal
}

func (c configKey) GetOptional() (string, bool) {
	val := configProvider.GetString(string(c))
	if val == "" {
		return val, false
	}
	return val, true
}

func (c configKey) GetStringMapString() map[string]string {
	return configProvider.GetStringMapString(string(c))
}

func (c configKey) GetStringSlice() []string {
	val := configProvider.GetStringSlice(string(c))
	if val == nil {
		return []string{}
	}
	return val
}
