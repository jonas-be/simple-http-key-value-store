package database

type Database map[string]string

func (database Database) Get(key string) string {
	return database[key]
}

func (database Database) Set(key string, value string) {
	database[key] = value
}

func (database Database) Delete(key string) {
	delete(database, key)
}

func (database Database) Contains(key string) bool {
	return database[key] != ""
}
