package mqtt

import (
	"bambulab-exporter/collector"
	"bambulab-exporter/config"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

func Start(reportReceiver func(report map[string]any)) <-chan error {
	errChan := make(chan error, 1)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", config.PrinterHost, 8883))
	opts.SetClientID("bambulab-exporter")
	opts.SetUsername("bblp")
	opts.SetPassword(config.PrinterAccessCode)
	opts.SetDefaultPublishHandler(onMessageDefault)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetMaxReconnectInterval(30 * time.Second)
	opts.OnConnect = onConnect(reportReceiver)
	opts.OnConnectionLost = onConnectionLost

	client := mqtt.NewClient(opts)

	go func() {
		tok := client.Connect()
		tok.Wait()
		if err := tok.Error(); err != nil {
			log.Warn().Err(err).Msg("Initial connection failed, will retry automatically")
			collector.MarkOffline()
		}
	}()

	return errChan
}

func onConnect(reportReceiver func(report map[string]any)) func(mqtt.Client) {
	return func(client mqtt.Client) {
		log.Info().Msg("Connected to MQTT broker")

		token := client.Subscribe(fmt.Sprintf("device/%s/report", config.PrinterSerial), 0, onMessage(reportReceiver))
		token.Wait()
		if err := token.Error(); err != nil {
			log.Error().Err(err).Msg("Failed to subscribe to MQTT topic")
			collector.MarkOffline()
		}
	}
}

func onConnectionLost(client mqtt.Client, err error) {
	log.Warn().Err(err).Msg("Connection lost, will attempt to reconnect")
	collector.MarkOffline()
}

func onMessageDefault(client mqtt.Client, msg mqtt.Message) {
	log.Debug().Str("payload", string(msg.Payload())).Msg("Received unexpected message")
}

func onMessage(reportReceiver func(report map[string]any)) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		log.Trace().Str("payload", string(payload)).Msg("Received message")

		var report map[string]any
		err := json.Unmarshal(payload, &report)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal message")
			return
		}

		reportReceiver(report)
	}
}
