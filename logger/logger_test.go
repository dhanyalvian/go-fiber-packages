//- logger/logger_test.go

package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		info       string
		message    string
		wantColor  string
	}{
		{
			name:       "Generic status code 0",
			statusCode: 0,
			info:       "INFO",
			message:    "This is a generic log",
			wantColor:  "\x1b[34m", // Blue
		},
		{
			name:       "Success status code 200",
			statusCode: 200,
			info:       "SUCCESS",
			message:    "Operation successful",
			wantColor:  "\x1b[32m", // Green
		},
		{
			name:       "Redirection status code 301",
			statusCode: 301,
			info:       "REDIRECT",
			message:    "Resource moved",
			wantColor:  "\x1b[33m", // Yellow
		},
		{
			name:       "Client error status code 404",
			statusCode: 404,
			info:       "CLIENT ERROR",
			message:    "Not found",
			wantColor:  "\x1b[33m", // Yellow
		},
		{
			name:       "Server error status code 500",
			statusCode: 500,
			info:       "SERVER ERROR",
			message:    "Internal server error",
			wantColor:  "\x1b[31m", // Red
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer log.SetOutput(nil)

			Logging(tt.statusCode, tt.info, tt.message)

			got := buf.String()
			if !strings.Contains(got, tt.info+": "+tt.wantColor+tt.message+"\x1b[0m") {
				t.Errorf("Logging() output = %q, want color sequence %q and message %q", got, tt.wantColor, tt.message)
			}
		})
	}
}
