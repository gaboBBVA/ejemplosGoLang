package main

import (
	"fmt"
)

/*
  NO ES NECESARIO EXPLICITAMENTE, COMO EN OTROS LENGUAJES, QUE SE IMPLEMENTA CIERTA
FUNCIONALIDAD (Java: NO implements)
*/


//Persona ...
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Nombre ... METHOD
func (p Persona) Nombre() string {
	return p.nombre
}

//Email ... METHOD
func (p Persona) Email() string {
	return p.email
}

//Moderador ...
type Moderador struct {
	Persona
	Foro string
}

//AbrirForo ... MODERADOR
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un Foro")
}

//CerrarForo ... MODERADOR
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar un Foro")
}

//Administrador ...
type Administrador struct {
	Persona
	Sección string
}

//CrearForo ... ADMINISTRADOR
func (a Administrador) CrearForo() {
	fmt.Println("Abrir un Foro")
}

//EliminarForo ... ADMINISTRADOR
func (a Administrador) EliminarForo() {
	fmt.Println("Cerrar un Foro")
}

//Presentarse ...
// 1) Persona
// 2) COMO INTERFACE
//func Presentarse(p Persona) {
//	fmt.Println("Nombre: ", p.Nombre())
//	fmt.Println("Email: ", p.Email())
//}

func Presentarse(p Usuario) {
	fmt.Println("Nombre: ", p.Nombre())
	fmt.Println("Email: ", p.Email())
}


//PresentarseA ...
func PresentarseA(a Administrador) {
 	fmt.Println("Nombre: ", a.Nombre())
 	fmt.Println("Email: ", a.Email())
}

//PresentarseM ...
func PresentarseM(m Moderador) {
 	fmt.Println("Nombre: ", m.Nombre())
 	fmt.Println("Email: ", m.Email())
}

//Usuario ...INTERFASE
type Usuario interface {
	Nombre() string
	Email() string
}

func main() {

	gabriel := Persona{"gabriel", "gabo@bbva.com", 10}
	Presentarse(gabriel)

	// SE REPITEN LAS DOS FUNCIONES CON EXACTAMENTE LA MISMA FUNCIONALIDAD
	//juan := Moderador{Persona{"Juan", "j@hmail.com", 46}, "Juegos"}
	//pedro := Administrador{Persona{"Pedro", "p@hmail.com", 25}, "PC"}
	//Presentarse(juan)
	//Presentarse(pedro)

	//  INTERFAZ USUARIO PARA MANIPULAR OBJETOS
	//var i Usuario
	//i = gabriel
	//fmt.Println("i: ", i)
	//fmt.Println("i: ", i.Email())
	//fmt.Println("i: ", i.Nombre())
	//
	//i = juan
	//fmt.Println("i: ", i)
	//fmt.Println("i: ", i.Email())
	//fmt.Println("i: ", i)
	//fmt.Println("i: ", i.CerrarForo()) // ERROR POR NO PERTENECER A LA INTERFAZ

}