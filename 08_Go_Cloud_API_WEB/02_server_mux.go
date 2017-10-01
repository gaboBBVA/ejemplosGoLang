package main

import (
	"net/http"
	"fmt"
)

func main(){
	mux := http.NewServeMux() // ENRUTADOR A UN HANDLER
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"<h1>Hola Paas Studio WEB con mux</h1>")
	})

	mux.HandleFunc("/paas", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"<h1>Hola Paas con mux desde /paas </h1>")
	})

	mux.HandleFunc("/aso", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"<h1>Hola ASO con mux desde /aso </h1>")
	})


	http.ListenAndServe(":8080",mux)
}
