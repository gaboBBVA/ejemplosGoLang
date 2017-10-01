package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
)

type mensaje struct {
	msj string
}

// FUNCION QUE DEBE IMPLEMENTAR LA INTERFAZ http.ServeHTTP
func (m mensaje) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msj)
}

func main () {
	/*
	Mensaje := mensaje {
		msj: "Hola otra vez",
	}
	*/

	Mensaje2 := mensaje {
		msj: "Paas, Hola otra vez",
	}


	Mensaje3 := mensaje {
		msj: "ASO",
	}


	mux := http.NewServeMux()
	// FileServer
	fs := http.FileServer(http.Dir("public"))


	mux.Handle("/", fs)

	mux.Handle("/paas", Mensaje2)

	mux.Handle("/aso", Mensaje3)

	//http.ListenAndServe(":8080", mux)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando en el puerto 8080 ...")
	server.ListenAndServe()
}