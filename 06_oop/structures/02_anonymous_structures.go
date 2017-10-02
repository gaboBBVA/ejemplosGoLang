package main

import (
	"fmt"
)

//Persona contiene los campos nombre y apellido
type Persona struct {
	Nombre   string
	Apellido string
}

//Estudiante contiene el campo persona y carrera
type Estudiante struct {
	Persona       // COMPONENTE ANONIMO SOLO EL TIPO - AUTOMATICAMENTE Estudiante TIENE LO DE Persona
	Carrera string
}

//Profesor tiene los campos Estudiante y Carrera
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	//Declaramos una variable del tipo Estudiante
	Gabriel := Estudiante{
		Persona{
			Nombre:   "Gabriel",
			Apellido: "Ramirez",
		},
		"MAC",
	}
	fmt.Println("Gabriel: ", Gabriel)

	//Accediendo a los campos
	fmt.Println("Nombre: ", Gabriel.Nombre)
	fmt.Println("Apellido: ", Gabriel.Apellido)
	fmt.Println("Carrera: ", Gabriel.Carrera)

	//Declaramos una clase de tipo Profesor
	pedro := Profesor{
		Estudiante{
			Persona{
				"Pedro",
				"Almonte",
			},
			"Cotabilidad",
		},
		"Informatica",
	}
	fmt.Println("Pedro: ", pedro)
	//
	//Accediendo a los campos
	fmt.Println("Nombre: ", pedro.Nombre)
	fmt.Println("Apellido: ", pedro.Apellido)
	fmt.Println("Carrera: ", pedro.Carrera)
	//
	fmt.Println("Carrera Estudiantes: ", pedro.Estudiante.Carrera)
	//
	//Declarando una variable de tipo Profesor
	manuel := Profesor{Estudiante{Persona{"Manuel", "Peralta"}, "Ing. Industrial"}, "Informatica"}
	fmt.Println("Manuel: ", manuel)
	//
	var chuy Profesor
	chuy.Nombre = "Jesus"
	chuy.Apellido = "Contreras"
	chuy.Estudiante.Carrera = "Informatica"
	chuy.Carrera = "Mercadotecnia"
	fmt.Println("José: ", chuy)

}