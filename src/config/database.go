package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetDatabase conexion al a base de datos
func GetDatabase() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "northwind"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
