package database

import (
	"sync"
)

type OutOfStorageError struct {
	msg string
}

func (e OutOfStorageError) Error() string {
	return e.msg
}

type InputError struct {
	msg string
}

func (e InputError) Error() string {
	return e.msg
}

type Database struct {
	mu   sync.Mutex
	Data map[string]string
}

func (db *Database) Get(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.Data[key]
}

func (db *Database) Set(key string, value string) error {
	if len(key) > 200 {
		return InputError{"key length exceeded"}
	}
	if len(value) > 200 {
		return InputError{"value length exceeded"}
	}
	db.mu.Lock()
	if len(db.Data) >= 50 {
		return OutOfStorageError{"database has already 50 entries"}
	}
	db.Data[key] = value
	db.mu.Unlock()
	return nil
}

func (db *Database) Delete(key string) {
	db.mu.Lock()
	delete(db.Data, key)
	db.mu.Unlock()
}

func (db *Database) Contains(key string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.Data[key] != ""
}
