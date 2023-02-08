package models

import (
	"log"
	database "newProject/Database"
)

type Pessoas struct {
	Id     int
	Nome   string
	Idade  int
	Cidade string
}

func NovaPessoa(nome string, idade int, cidade string) {

	db := database.ConectarDatabase()

	insertDados, err := db.Prepare("insert into Clientes(nome, idade, cidade) values($1,$2,$3)")

	if err != nil {
		log.Println("ERRO: erro na inserção de dados na tabela Clientes")
		panic(err.Error())
	}

	insertDados.Exec(nome, idade, cidade)
	defer db.Close()

}
