package main

import (
	"time"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
	"log"
)

// API REST - ENTIDAD Nota
type Nota struct {
	Titulo string `json:"titulo"` 			// ANOTACION DE CAMPOS DE ESTRUCTURAS
	Descripcion string `json:"descripcion"`
	CreadaElDia time.Time `json:"creada_el_dia"`
}

var datosNotas = make (map[string]Nota)
var id int

func main (){
	// ENRUTADOR DE GORILLA MUX
	gorilla_router := mux.NewRouter().StrictSlash(false)

	// MANEJADORES POR VERBOS HTTP
	gorilla_router.HandleFunc("/api/notes",GetNoteHandler).Methods("GET")
	gorilla_router.HandleFunc("/api/notes",PostNoteHandler).Methods("POST")
	gorilla_router.HandleFunc("/api/notes/{id}",PutNoteHandler).Methods("PUT")
	gorilla_router.HandleFunc("/api/notes/{id}",DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr: ":8080",
		Handler: gorilla_router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando en localhost puerto 8080 ...")
	server.ListenAndServe()



}

// GetNoteHandler - GET - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request){
	// SLICE DE Notas
	var notas []Nota
	for _,valor := range datosNotas {
		notas = append(notas,valor)
	}
	// Set ESTABLECE CABECERAS HTTP
	w.Header().Set("Content-Type","application/json")

	// PASAR LOS DATOS AL FORMATO JSON CON Marshall
	j, err := json.Marshal(notas)
	if err != nil {
		panic (err)
	}
	// ESCRIBIR LA RESPUESTA HTTP PARA EL USUARIO
	w.WriteHeader(http.StatusOK)

	// CONTENIDO  Y RESPUESTA EN FORMATO JSON
	w.Write(j)

}

// PostNoteHandler - POST - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request){
	// SLICE DE Notas
	var nota Nota
	// DECODIFICADOR PARA EL DATO DE ENTRADA Y VERIFICAR QUE EL JSON ENVIADO ES CORRECTO
	err:=json.NewDecoder(r.Body).Decode(&nota)
	if err != nil {
		panic(err)
	}
	// SE AÑADE EL TIME STAMP A LA Nota
	nota.CreadaElDia = time.Now()
	id++
	k := strconv.Itoa(id)
	datosNotas[k] = nota

	// SE PREPARA LA RESPUESTA AL CLIENTE
	// Set ESTABLECE CABECERAS HTTP
	w.Header().Set("Content-Type","application/json")

	// PASAR LOS DATOS AL FORMATO JSON CON Marshall
	j, err := json.Marshal(nota)
	if err != nil {
		panic (err)
	}
	// ESCRIBIR LA RESPUESTA HTTP PARA EL USUARIO
	w.WriteHeader(http.StatusCreated)

	// CONTENIDO  Y RESPUESTA EN FORMATO JSON
	w.Write(j)
}

// PutNoteHandler - PUT - /api/notes
func PutNoteHandler(w http.ResponseWriter, r *http.Request){
	// SE RECUPERAN LOS PARAMETROS, EN ESTE CASO EL id
	// DEVUELVE EN UN MAP DE STRING CUYO INDICE ES EL NOMBRE DE LA VARIABLE
	vars := mux.Vars(r)
	k := vars["id"]

	// SE OBTIENEN LOS DATOS QUE SE VAN A ACTUALIZAR
	var notaUpdate Nota
	err := json.NewDecoder(r.Body).Decode(&notaUpdate)
	if err != nil {
		panic(err)
	}

	// SE REVISA SI EXISTE LA NOTA
	if nota, ok := datosNotas[k];ok {
		// SE RECUPERA EL TIMESTAMP DE LA NOTA A ACTUALIZAR
		notaUpdate.CreadaElDia = nota.CreadaElDia
		// SE BORRA LA NOTA ANTERIOR
		delete (datosNotas,k)
		// SE ACTUALIZA EL REGISTRO CON LOS DATOS NUEVOS
		datosNotas[k] = notaUpdate
	} else {
		log.Printf("No se encontro la nota con el id es: %d",k )
	}

	// SE MANDA LA RESPUESTA AL CLIENTE
	w.WriteHeader(http.StatusNoContent)

}

// DeleteNoteHandler - DELETE - /api/notes
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request){
	// SE RECUPERAN LOS PARAMETROS, EN ESTE CASO EL id
	// DEVUELVE EN UN MAP DE STRING CUYO INDICE ES EL NOMBRE DE LA VARIABLE
	vars := mux.Vars(r)
	k := vars["id"]

	// SE REVISA SI EXISTE LA NOTA
	if _, ok := datosNotas[k];ok {
		delete (datosNotas,k)
	} else {
		log.Printf("No se encontro la nota con el id es: %d",k )
	}

	// SE MANDA LA RESPUESTA AL CLIENTE
	w.WriteHeader(http.StatusNoContent)

}