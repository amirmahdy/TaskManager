package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBConn              string        `mapstructure:"DB_CONN"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(p string) (config Config, err error) {
	viper.AddConfigPath(p)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
