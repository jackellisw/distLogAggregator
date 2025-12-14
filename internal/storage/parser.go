package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackellisw/distLogAggregator.git/internal/models"
)

func ParseLog(rawLog string) (models.LogEntry, error) {
	parts := strings.SplitN(rawLog, " ", 4)

	serviceAndMessage := strings.Split(parts[3], ":")
	serviceName := serviceAndMessage[0]
	message := strings.Trim(serviceAndMessage[1], " ")

	logLevel := strings.Trim(parts[2], "[]")

	timeStamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
	if err != nil {
		return models.LogEntry{}, fmt.Errorf("Error parsing log: %s", err)
	}

	parseLog := models.LogEntry{
		TimeStamp:   timeStamp,
		LogLevel:    logLevel,
		ServiceName: serviceName,
		Message:     message,
	}

	return parseLog, nil
}
