package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitConfig() Config {
	cfg, err := readConfigFromEnv()
	if err != nil {
		panic("cannot initialize .env config")
	}

	return cfg
}

func readConfigFromEnv() (Config, error) {
	var cfg Config

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("error while reading .env config file")
		return Config{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to parsing .env config")
		return Config{}, err
	}

	return cfg, nil

}
