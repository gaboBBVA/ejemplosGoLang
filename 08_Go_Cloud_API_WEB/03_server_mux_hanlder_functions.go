package main

import (
	"net/http"
	"fmt"
)

func prueba (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"<h1>Hola Paas Studio WEB con mux y function handlers</h1>")
}

func paas (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"<h1>Hola Paas desde function handler </h1>")
}

func aso (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"<h1>Hola ASO desde function handler </h1>")
}



func main(){
	mux := http.NewServeMux() // ENRUTADOR A UN HANDLER

	mux.HandleFunc("/", prueba)

	mux.HandleFunc("/paas", paas)

	mux.HandleFunc("/aso", aso)

	http.ListenAndServe(":8080",mux)
}
