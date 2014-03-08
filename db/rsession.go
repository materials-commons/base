package db

import (
	r "github.com/dancannon/gorethink"
)

var dbAddress = ""
var dbName = ""

func SetAddress(address string) {
	dbAddress = address
}

func SetDatabase(db string) {
	dbName = db
}

func RSession() (*r.Session, error) {
	return r.Connect(map[string]interface{}{
		"address":  dbAddress,
		"database": dbName,
	})
}
