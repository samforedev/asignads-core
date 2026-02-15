package asigna_multitenancy

import (
	"sync"

	"gorm.io/gorm"
)

type ConnectionRegistry struct {
	mu          sync.RWMutex
	connections map[string]*gorm.DB
	dsnCache    map[string]string
}

func NewConnectionRegistry() *ConnectionRegistry {
	return &ConnectionRegistry{
		connections: make(map[string]*gorm.DB),
		dsnCache:    make(map[string]string),
	}
}

func (r *ConnectionRegistry) Get(id string) (*gorm.DB, string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	conn, exists := r.connections[id]
	dsn := r.dsnCache[id]
	return conn, dsn, exists
}

func (r *ConnectionRegistry) Set(id string, dsn string, db *gorm.DB) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.dsnCache[id] = dsn
	r.connections[id] = db
}
