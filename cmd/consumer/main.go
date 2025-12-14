package main

import (
	"fmt"

	"github.com/jackellisw/distLogAggregator.git/internal/parser"
	"github.com/jackellisw/distLogAggregator.git/internal/storage"
)

func main() {
	testLogs := []string{
		"2024-12-14 10:30:45 [ERROR] UserService: Failed to connect to database",
		"2024-12-14 10:34:45 [ERROR] UserService: Failed to connect to database",
		"2024-12-14 10:39:45 [ERROR] UserService: Failed to connect to database",
	}

	store := parser.LogStore{}

	for _, e := range testLogs {
		parsed, err := storage.ParseLog(e)
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
