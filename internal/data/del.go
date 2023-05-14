package data

func (db *Database) Del(key string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	_, exists := db.data[key]
	if exists {
		delete(db.data, key)
	}

	return exists
}
