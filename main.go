package main

import (
	"net/http"
	"newProject/Routes"
)

func main() {

	routes.Routes()
	http.ListenAndServe(":8000", nil)

}
