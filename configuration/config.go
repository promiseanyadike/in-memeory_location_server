package configuration

import (
	"os"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel    int
	ApiListener string
}

func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.LogLevel, v.Min(0), v.Max(5)),
		v.Field(&c.ApiListener),
	)
}

func InitConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("")
	}
	viper.AutomaticEnv()
	c := new(Config)
	c.LogLevel = viper.GetInt("LOG_LEVEL")
	apiListener := viper.GetString("HISTORY_SERVER_LISTEN_ADDR")
	if len(apiListener) > 0 {
		c.ApiListener = apiListener
	} else {
		c.ApiListener = ":8080"
	}

	if err := c.Validate(); err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(-1)
	}
	return c
}
