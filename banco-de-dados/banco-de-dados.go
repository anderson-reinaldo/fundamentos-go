package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "reinaldo:rei%RDXxdr%@/devbook?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão aberta")

	users, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

	defer users.Close()
}
