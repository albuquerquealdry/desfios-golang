package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de mysql
)

func Connect() (*sql.DB, error) {
	stringConnectDefault := "golang:golang@123Apççs%#@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConnectDefault)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
