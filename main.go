package main

import (
	"RestGoChi/database"
	"fmt"

	// "net/http"

	// "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := database.InitDB()
	defer dbConnection.Close()

	fmt.Println(dbConnection)
}
