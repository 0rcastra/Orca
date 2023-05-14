package data

func Get(db *Database, key string) (interface{}, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.data[key]

	return value, exists
}
