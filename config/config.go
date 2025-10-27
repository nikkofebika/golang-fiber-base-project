package config

import (
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort    string `mapstructure:"APP_PORT"`
	AppEnv     string `mapstructure:"APP_ENV"`
	AppPrefork bool   `mapstructure:"APP_PREFORK"`

	LogFormat   string `mapstructure:"LOG_FORMAT"`
	LogFilePath string `mapstructure:"LOG_FILE_PATH"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	JWTSecret           string        `mapstructure:"JWT_SECRET"`
	JWTRefreshSecret    string        `mapstructure:"JWT_REFRESH_SECRET"`
	JWTExpiresIn        time.Duration `mapstructure:"JWT_EXPIRES_IN"`
	JWTRefreshExpiresIn time.Duration `mapstructure:"JWT_REFRESH_EXPIRES_IN"`
}

func LoadAppConfig() (appConfig AppConfig) {
	viper.SetConfigFile(".env")

	var err error
	if err = viper.ReadInConfig(); err != nil {
		panic(".env failed to loaded")
	}

	if err = viper.Unmarshal(&appConfig); err != nil {
		panic(".env failed to loaded")
	}

	return
}
