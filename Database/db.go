package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectarDatabase() *sql.DB {

	conexao := "user=postgres dbname=postgres password=98911 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Println("Erro na tentativa de conexão DataBase")
		panic(err.Error())
	}
	return db
}
