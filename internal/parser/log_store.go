package parser

import (
	"sync"

	"github.com/jackellisw/distLogAggregator.git/internal/models"
)

type LogStore struct {
	mu   sync.Mutex
	logs []models.LogEntry
}

func (ls *LogStore) Add(entry models.LogEntry) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.logs = append(ls.logs, entry)
}

func (ls *LogStore) GetAll() []models.LogEntry {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	return ls.logs
}
