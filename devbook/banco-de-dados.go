package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnectDefault := "golang:golang@123Apççs%#@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConnectDefault)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Open Connection")

	lines, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	defer lines.Close()
	fmt.Println(lines)

}
