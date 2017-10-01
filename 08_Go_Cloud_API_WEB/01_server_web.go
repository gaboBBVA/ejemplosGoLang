package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"<h1>Hola Paas Studio WEB </h1>")
	})
	http.ListenAndServe(":8080",nil)
}
