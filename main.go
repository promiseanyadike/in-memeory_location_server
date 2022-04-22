package main

import (
	"context"
	"in-memory_location_server/configuration"
	"in-memory_location_server/logger"
	"in-memory_location_server/route"
	"in-memory_location_server/service"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Init config
	cf := configuration.InitConfig()
	// Init logger
	logger.InitPretty(zerolog.Level(cf.LogLevel))

	repo := service.NewRepo()
	locationService := service.NewLocation(&repo)

	// Init handler
	handle := route.NewHandler(&locationService)

	srv := &http.Server{Addr: cf.ApiListener, Handler: handle}
	log.Info().Msgf("Start service on http://%s", cf.ApiListener)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info().Msgf("Service - listen: %s", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		var once sync.Once
		for range signalChan {
			once.Do(func() {
				log.Warn().Msgf("Service received a shutdown signal...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := srv.Shutdown(ctx); err != nil {
					log.Error().Err(err).Msg("Error")
				}
				log.Info().Msg("Service - Received an interrupt closing connection...")
				log.Warn().Msgf("Service stopped successfully")
				cleanupDone <- true
			})
		}
	}()
	log.Warn().Msgf("Service started successfully")
	<-cleanupDone
}
