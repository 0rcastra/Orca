package data

import (
	"fmt"
)

func (db *Database) Incr(key string) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	value, exists := db.data[key]
	if !exists {
		db.data[key] = 1
		return 1, nil
	}

	intValue, ok := value.(int)
	if !ok {
		return 0, fmt.Errorf("value for key '%s' is not an integer", key)
	}

	intValue++
	db.data[key] = intValue

	return intValue, nil
}
