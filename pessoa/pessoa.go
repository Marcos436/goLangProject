package pessoa

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

var p pessoa

func Falar() {

	fmt.Println("Qual e o seu nome?: ")
	fmt.Scan(&p.nome)

	fmt.Println("Qual e a sua Idade: ")
	fmt.Scan(&p.idade)

	fmt.Print("Hello World!\n")
	fmt.Println("Eu sou o", p.nome, "tenho:", p.idade, "anos")

}
