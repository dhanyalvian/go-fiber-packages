//- logger/logger.go

package logger

import "log"

func Logging(statusCode int, info string, message string) {
	var colorStart string
	colorEnd := "\x1b[0m"

	switch {
	case statusCode == 0:
		colorStart = "\x1b[34m" // Blue for generic
	case statusCode >= 200 && statusCode < 300:
		colorStart = "\x1b[32m" // Green for success
	case statusCode >= 300 && statusCode < 500:
		colorStart = "\x1b[33m" // Yellow for redirection/client errors
	default:
		colorStart = "\x1b[31m" // Red for server errors
	}

	log.Printf("%s: %s%s%s", info, colorStart, message, colorEnd)
}
