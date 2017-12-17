package database

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
)

var db *storm.DB

//Open database connection
func OpenDB(path string) error {
	newDb, err := storm.Open(path, storm.Codec(gob.Codec))
	if err != nil {
		return err
	}
	db = newDb

	return nil
}

//Get database connection
func GetDB() *storm.DB {
	return db
}
