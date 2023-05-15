package data

import (
	"fmt"
	"strconv"
)

func (db *Database) Incr(key string) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	value, exists := db.data[key]
	if !exists {
		db.data[key] = "1"
		return 1, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("value for key '%s' is not a valid integer: %w", key, err)
	}

	intValue++
	db.data[key] = strconv.Itoa(intValue)

	return intValue, nil
}
