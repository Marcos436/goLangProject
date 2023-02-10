package pessoa

import (
	"fmt"
	
)

type pessoa struct {
	Id int 
	Nome  string
	Idade int
	Cidade string
	
}

var p pessoa

func Falar() {


	fmt.Println("Qual e o seu nome?: ")
	fmt.Scan(&p.Nome)

	fmt.Println("Qual e a sua Idade: ")
	fmt.Scan(&p.Idade)

	fmt.Print("Hello World!\n")
	fmt.Println("Eu sou o", p.Nome, "tenho:", p.Idade, "anos")

}
