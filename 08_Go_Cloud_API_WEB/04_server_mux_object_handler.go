package main

import (
	"net/http"
	"fmt"
)

type mensaje struct {
	msj string
}

// FUNCION QUE DEBE IMPLEMENTAR LA INTERFAZ http.ServeHTTP
func (m mensaje) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msj)
}

func main () {
	Mensaje := mensaje {
		msj: "Hola otra vez",
	}

	Mensaje2 := mensaje {
		msj: "Paas, Hola otra vez",
	}


	Mensaje3 := mensaje {
		msj: "ASO",
	}


	mux := http.NewServeMux()

	mux.Handle("/", Mensaje)

	mux.Handle("/paas", Mensaje2)

	mux.Handle("/aso", Mensaje3)

	http.ListenAndServe(":8080", mux)
}