package database

import "github.com/asdine/storm"

var db *storm.DB

func OpenDB(path string) error {
	newDb, err := storm.Open(path)
	if err != nil {
		return err
	}
	db = newDb

	return nil
}

func GetDB() *storm.DB {
	return db
}
