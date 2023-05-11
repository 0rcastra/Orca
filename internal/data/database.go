package data

import (
    "sync"
)

// Database Struct
type Database struct {
    data  map[string]string
    mutex sync.RWMutex
}

// creates a new instance of the Database.
func NewDatabase() *Database {
    return &Database{
        data: make(map[string]string),
    }
}
