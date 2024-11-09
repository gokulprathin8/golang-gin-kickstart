package database

import "database/sql"

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(map[string]interface{}{
			"info": "failed to open database",
			"err":  err,
		})
	}
}
