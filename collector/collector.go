package collector

import (
	"bambulab-exporter/config"
	"bambulab-exporter/metrics"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// seenUnknownKeys tracks which unknown keys have already been logged
var seenUnknownKeys = make(map[string]bool)
var seenUnknownKeysMutex sync.RWMutex

// offlineTimer tracks when to clear metrics after prolonged disconnection
var offlineTimer *time.Timer
var offlineTimerMutex sync.Mutex

const offlineGracePeriod = 30 * time.Second

// isNewUnknownField checks if an unknown field has been seen before.
// It returns true if the field has not been seen before, and false otherwise.
// It also marks the field as seen.
func isNewUnknownField(reportType, reportKey string) bool {
	key := reportType + "|" + reportKey

	seenUnknownKeysMutex.RLock()
	alreadySeen := seenUnknownKeys[key]
	seenUnknownKeysMutex.RUnlock()
	if alreadySeen {
		return false
	}

	seenUnknownKeysMutex.Lock()
	alreadySeen = seenUnknownKeys[key] // check again in case another goroutine saw this key while we were waiting for the lock
	seenUnknownKeys[key] = true
	seenUnknownKeysMutex.Unlock()
	return !alreadySeen
}

func ReceiveReport(report map[string]any) {
	metrics.PrinterOnline.WithLabelValues(config.PrinterName).Set(1)

	// Cancel offline timer since we received a report
	offlineTimerMutex.Lock()
	if offlineTimer != nil {
		offlineTimer.Stop()
		offlineTimer = nil
	}
	offlineTimerMutex.Unlock()

	for key, value := range report {
		switch key {
		case "print":
			printReport, ok := value.(map[string]any)
			if !ok {
				log.Error().Interface("value", value).Msg("Invalid print report")
				continue
			}
			receivePrintReport(printReport)
		default:
			if isNewUnknownField("", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in report")
			}
		}
	}
}

func MarkOffline() {
	metrics.PrinterOnline.WithLabelValues(config.PrinterName).Set(0)

	// Start a timer to clear metrics after grace period
	// This prevents clearing metrics during brief connection drops
	offlineTimerMutex.Lock()
	if offlineTimer != nil {
		offlineTimer.Stop()
	}
	offlineTimer = time.AfterFunc(offlineGracePeriod, func() {
		log.Info().Msg("Printer offline for grace period, clearing metrics")
		metrics.ClearAllMetrics(config.PrinterName)
	})
	offlineTimerMutex.Unlock()
}

// safeFloat64 converts a value to float64, returning 0 if conversion fails
func safeFloat64(value any) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
		// Handle special cases like "-47dBm"
		if strings.HasSuffix(v, "dBm") {
			if f, err := strconv.ParseFloat(strings.TrimSuffix(v, "dBm"), 64); err == nil {
				return f
			}
		}
		// Handle percentage strings like "100"
		if strings.HasSuffix(v, "%") {
			if f, err := strconv.ParseFloat(strings.TrimSuffix(v, "%"), 64); err == nil {
				return f
			}
		}
		// Handle time strings like "8" (hours)
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}

// safeBool converts a value to bool, returning false if conversion fails
func safeBool(value any) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v == "true" || v == "1" || v == "enable" || v == "on"
	}
	return false
}

// safeString converts a value to string, returning empty string if conversion fails
func safeString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	}
	return ""
}
