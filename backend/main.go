package main

import (
	"github.com/rs/zerolog/log"
	"hao123/api"
	"hao123/db"
	"hao123/util"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("Load Config Failed")
	}

	connect, err := db.Connect(config.DatabaseURL, config.DBName, config.DBUsername, config.DBPassword, config.Source, 3) // Use a test database
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer func() {
		if err := connect.Disconnect(connect); err != nil {
			log.Printf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	queries := db.NewMongoQueries(connect.Database)
	store := db.NewMongoStore(connect.Client, queries)

	runHttpServer(store, config)

	go listenForShutdown()

}

func runHttpServer(store db.Store, config util.Config) {

	s, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	err = s.Start(config.HttpServerAddr)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot Start Sever..")
	}
}

func listenForShutdown() {
	var wg *sync.WaitGroup

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	shutdown(wg)
	os.Exit(0)
}
func shutdown(wg *sync.WaitGroup) {

	log.Info().Msg("would run clean up tasks...")

	wg.Wait()

	log.Info().Msg("closing channels and shutting down applications...")

}
