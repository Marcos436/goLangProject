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
	Senha int
}

func NovaPessoa(nome, cidade string, idade,senha int) {

	db := database.ConectarDatabase()

	insertDados, err := db.Prepare("insert into Clientes(nome, idade, cidade,senha) values($1,$2,$3,$4)")

	if err != nil {
		log.Println("ERRO: erro na inserção de dados na tabela Clientes")
		panic(err.Error())
	}

	insertDados.Exec(nome, idade, cidade,senha)
	defer db.Close()

}
