package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env             string `mapstruct:"ENV"`
	Port            string `mapstructure:"PORT"`
	DBURL           string `mapstructure:"DATABASE_URL"`
	DBName          string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, nil
		}
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}