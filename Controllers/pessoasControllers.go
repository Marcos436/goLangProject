package controllers

import (
	"log"
	"net/http"
	models "newProject/Models"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/page/*.html"))

func Login(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "Login", nil)

}

func Cadastro(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "cadastro", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		idade := r.FormValue("idade")
		cidade := r.FormValue("cidade")
		senha := r.FormValue("senha")

		idadeconvert, err := strconv.Atoi(idade)

		if err != nil {
			log.Println("ERRO: erro na conversão preco")
			panic(err.Error())
		}

		senhaconvert, err := strconv.Atoi(senha)

		if err != nil {

			panic(err.Error())
		}

		models.NovaPessoa(nome, cidade, idadeconvert, senhaconvert)
	}
	http.Redirect(w, r, "Login.html", 301)

}
