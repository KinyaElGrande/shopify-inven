package data

import (
	"database/sql"
	"mumbi/inven-logistics/config"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := config.GoDotEnvVariable("DB_USER")
	dbPass := config.GoDotEnvVariable("DB_PASS")
	dbName := config.GoDotEnvVariable("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
