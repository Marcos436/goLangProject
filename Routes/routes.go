package routes

import (
	"net/http"
	"newProject/Controllers"
)
func Routes() {

	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/cadastro", controllers.Cadastro)
	http.HandleFunc("/insert", controllers.Insert)

}