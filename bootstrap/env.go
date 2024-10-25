package bootstrap

import (
	"blackwebfilter/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Env struct {
	AppEnv             string `mapstructure:"APP_ENV"`
	ServerAddress      string `mapstructure:"SERVER_ADDRESS"`
	LimitRequests      int    `mapstructure:"LIMIT_REQUESTS"`
	BlackDomainsPath   string `mapstructure:"BLACK_DOMAINS_PATH"`   // blackweblist file path
	SensitiveWordsPath string `mapstructure:"SENSITIVE_WORDS_PATH"` // sensitiveword file path
	ZapLevel           string `mapstructure:"ZAP_LEVEL"`
	ZapPath            string `mapstructure:"ZAP_PATH"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
}

func NewEnv() *Env {
	env := &Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Can't find the file .env : ", zap.Error(err))
	}
	err = viper.Unmarshal(env)
	if err != nil {
		logger.Error("Environment can't be loaded: ", zap.Error(err))
	}
	if env.AppEnv == "development" {
		logger.Infof("The App is running in development env")
	} else if env.AppEnv == "production" {
		logger.Infof("The App is running in production env")
	}
	return env
}
