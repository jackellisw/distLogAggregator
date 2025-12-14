package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	TimeStamp   time.Time
	LogLevel    string
	ServiceName string
	Message     string
}

type LogStore struct {
	mu   sync.Mutex
	logs []LogEntry
}

func (ls *LogStore) Add(entry LogEntry) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.logs = append(ls.logs, entry)
}

func (ls *LogStore) GetAll() []LogEntry {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	return ls.logs
}

// parse the raw log from producer and return a log entry
func ParseLog(rawLog string) (LogEntry, error) {
	parts := strings.SplitN(rawLog, " ", 4)

	serviceAndMessage := strings.Split(parts[3], ":")
	serviceName := serviceAndMessage[0]
	message := strings.Trim(serviceAndMessage[1], " ")

	logLevel := strings.Trim(parts[2], "[]")

	timeStamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
	if err != nil {
		return LogEntry{}, fmt.Errorf("Error parsing log: %s", err)
	}

	parseLog := LogEntry{
		TimeStamp:   timeStamp,
		LogLevel:    logLevel,
		ServiceName: serviceName,
		Message:     message,
	}

	return parseLog, nil
}

func main() {
	testLogs := []string{
		"2024-12-14 10:30:45 [ERROR] UserService: Failed to connect to database",
		"2024-12-14 10:34:45 [ERROR] UserService: Failed to connect to database",
		"2024-12-14 10:39:45 [ERROR] UserService: Failed to connect to database",
	}

	store := LogStore{}

	for _, e := range testLogs {
		parsed, err := ParseLog(e)
		if err != nil {
			fmt.Println("error parsing log:", err)
		}

		store.Add(parsed)
	}

	allLogs := store.GetAll()
	for i, e := range allLogs {
		formatedLog := fmt.Sprintf(
			"%s %d %s @ %s: %s",
			e.LogLevel,
			i,
			e.ServiceName,
			e.TimeStamp,
			e.Message,
		)
		fmt.Println(formatedLog)

	}
}
