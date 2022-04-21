package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const logTimeFormat = "02/01/2006 - 15:04:05.999 MST"

var once sync.Once

func Init(level zerolog.Level) {
	once.Do(func() {
		zerolog.SetGlobalLevel(level)
		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	})
}

func InitPretty(level zerolog.Level) {
	once.Do(func() {
		zerolog.SetGlobalLevel(level)
		log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: logTimeFormat})
	})
}
