package database

import "database/sql"

func InitDB() *sql.DB {
	dbConnection, err := sql.Open("mysql", "root:HOLAcomoestas88--@/nortwind")
	if err != nil {
		panic(err.Error())
	}
	return dbConnection

}
