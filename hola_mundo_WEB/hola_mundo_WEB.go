package main

import (
	"net/http"
	_ "net/http/pprof"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/",handler)
	error:= http.ListenAndServe(":6060", nil)
	if error != nil{
		log.Fatal("Error en el servidor:", error)
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hola mundo Go WEB para PaaS Studio")
}