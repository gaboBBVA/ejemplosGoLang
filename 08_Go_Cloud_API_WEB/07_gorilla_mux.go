package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"fmt"
)

// LAS FUNCIONES DEBEN LLEVAR LA FIRMA DE UN HandlerFunc
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Respondiendo desde el metodo GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Respondiendo desde el metodo POST")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Respondiendo desde el metodo PUT")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Respondiendo desde el metodo DELETE")
}

func main () {
	gorilla_router := mux.NewRouter().StrictSlash(false) // DISCRIMINAR CON O SIN SLASH

	gorilla_router.HandleFunc("/api/users", GetUsers).Methods("GET")
	gorilla_router.HandleFunc("/api/users", PostUsers).Methods("POST")
	gorilla_router.HandleFunc("/api/users", PutUsers).Methods("PUT")
	gorilla_router.HandleFunc("/api/users", DeleteUsers).Methods("DELETE")

	servidor := &http.Server{
		Addr: ":8080",
		Handler: gorilla_router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando en el puerto 8080 ...")
	servidor.ListenAndServe()

}

