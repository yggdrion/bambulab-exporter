package main

import (
	"bambulab-exporter/api"
	"bambulab-exporter/collector"
	"bambulab-exporter/config"
	"bambulab-exporter/logging"
	"bambulab-exporter/mqtt"

	"github.com/rs/zerolog/log"
)

func main() {
	logging.Setup()
	err := config.LoadEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	apiErrChan := api.Start()
	mqttErrChan := mqtt.Start(collector.ReceiveReport)

	log.Info().Msg("Bambulab exporter started")

	select {
	case err := <-apiErrChan:
		log.Fatal().Err(err).Msg("API error")
	case err := <-mqttErrChan:
		log.Fatal().Err(err).Msg("MQTT error")
	}
}
