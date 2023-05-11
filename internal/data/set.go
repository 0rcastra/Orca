package data

func (db *Database) Set(key, value string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data[key] = value
}
