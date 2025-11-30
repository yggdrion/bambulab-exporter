package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnv() (err error) {
	if err := godotenv.Load(); err != nil {
		log.Debug().Err(err).Msg("No .env file found, using environment variables")
	}
	return loadConfig()
}

var (
	PrinterHost       = ""
	PrinterSerial     = ""
	PrinterAccessCode = ""
	PrinterName       = ""
)

func loadConfig() error {
	PrinterHost = os.Getenv("BAMBULAB_EXPORTER_PRINTER_HOST")
	PrinterSerial = os.Getenv("BAMBULAB_EXPORTER_PRINTER_SERIAL")
	PrinterAccessCode = os.Getenv("BAMBULAB_EXPORTER_PRINTER_ACCESS_CODE")
	PrinterName = os.Getenv("BAMBULAB_EXPORTER_PRINTER_NAME")

	if PrinterHost == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_HOST is not set")
	}
	if PrinterSerial == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_SERIAL is not set")
	}
	if PrinterAccessCode == "" {
		return fmt.Errorf("BAMBULAB_EXPORTER_PRINTER_ACCESS_CODE is not set")
	}
	if PrinterName == "" {
		PrinterName = PrinterSerial
	}
	return nil
}
