package database

import "database/sql"

// InitDB comentario
func InitDB() *sql.DB {
	dbConnection, err := sql.Open("mysql", "root:HOLAcomoestas88--@/northwind")
	if err != nil {
		panic(err.Error())
	}
	return dbConnection

}
