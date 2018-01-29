package memory

import (
	"murrpy/model"
	"murrpy/store"
	"sync"
)

// Memory implements store
type Memory struct {
	db map[string]*model.Media
	mu sync.Mutex
}

// Get returns a media file
func (mem *Memory) Get(hash string) *model.Media {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	return mem.db[hash]
}

// Set a media object on the store
func (mem *Memory) Set(m *model.Media) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	if mem.db == nil {
		mem.db = make(map[string]*model.Media)
	}
	mem.db[m.Hash] = m
}

// New returns a memory store that implements the store interface
func New() store.Store {
	return &Memory{}
}
