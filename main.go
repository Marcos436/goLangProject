package main

import (
	"fmt"
	"newProject/pessoa"
)

var pessoas = pessoa.Falar()
var idade = pessoa.Idade()

func main() {

	fmt.Println("Qual e o seu nome?: ")
	fmt.Scan(&pessoas)

	fmt.Println("Qual e a sua Idade: ")
	fmt.Scan(&idade)

	fmt.Println("Hello World")
	fmt.Println("Eu sou o", pessoas, "tenho:", idade, "anos")
}
