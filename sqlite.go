package main

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type SqliteCache struct {
	db *sql.DB
}

func NewSqlCacheService() (Cache, error) {
	if db == nil {
		err := initSqliteCache()
		if err != nil {
			return &SqliteCache{}, err
		}
	}
	return &SqliteCache{db: db}, nil
}

func (s *SqliteCache) Name() string {
	return "Sqlite"
}

func initSqliteCache() error {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return errors.New("cannot open an SQLite memory database => " + err.Error())
	}
	_, err = db.Exec("CREATE TABLE cache (key string, value string);")
	if err != nil {
		return errors.New("cannot create schema => " + err.Error())
	}
	return nil
}

func (s *SqliteCache) Get(key string) (interface{}, error) {
	value, err := s.getFromCache(key)
	if err != nil && err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", errors.New("cannot scan addition => " + err.Error())
	}
	return value, nil
}

func (s *SqliteCache) Set(key string, data interface{}) error {
	if key == "" {
		return errors.New("Key is required")
	}
	_, err := s.getFromCache(key)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err != nil && err == sql.ErrNoRows {
		_, err = s.db.Exec("insert into cache (key, value) values (?,?)", key, data)
		return err
	}

	_, err = s.db.Exec("update cache set value = ? where key=?", data, key)
	return err
}

func (s *SqliteCache) getFromCache(key string) (string, error) {
	row := s.db.QueryRow("SELECT value FROM cache where key = ?", key)
	if row.Err() != nil {
		return "", row.Err()
	}
	value := ""
	err := row.Scan(&value)
	return value, err
}
