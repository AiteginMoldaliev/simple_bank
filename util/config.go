package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Dbdriver             string        `mapstructure:"DB_DRIVER"`
	Dbsource             string        `mapstructure:"DB_SOURCE"`
	GinServerAddress     string        `mapstructure:"GIN_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccesTokenDuration   time.Duration `mapstructure:"ACCESs_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
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
