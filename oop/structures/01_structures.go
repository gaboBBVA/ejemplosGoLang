package main

import (
	"fmt"
)

//Structures
// si dos tipos se pueden comparar entonces se pueden usar como indices en una estructura map

//Persona : estructura que representa a una persona
type Persona struct {
	Nombre string
	Edad   int
}

//func Mayor que determina cual es mayor y la diferencia de edad
func Mayor(p1, p2 Persona) (Persona, int) {
	if p1.Edad > p2.Edad {
		return p1, p1.Edad - p2.Edad
	}
	return p2, p2.Edad - p1.Edad
}

func main() {
	//Declara una variable de tipo Perona #1 por separado
	 var p Persona
	 p.Nombre = "Ramon"
	 p.Edad = 29
	 fmt.Println("Estructura p de tipo Persona: ", p)
	 fmt.Println("Nombre p: ", p.Nombre)
	 fmt.Println("Edad p: ", p.Edad)

	 //Declara e inicializa una variable de tipo Persona #2 en una misma sentencia explícitamente
	 p2 := Persona{Nombre: "Guadalupe", Edad: 10}
	 fmt.Println("Nombre p2: ", p2.Nombre)
	 fmt.Println("Edad p2: ", p2.Edad)

	 // Declara e inicializa una variable de tipo Persona #3 en una misma sentencia implícitamente
	 p3 := Persona{"Miguel", 18}
	 fmt.Println("Nombre p3: ", p3.Nombre)
	 fmt.Println("Edad p3: ", p3.Edad)

	// Inicializamos 3 variables de tipo Persona
	hugo := Persona{"hugo", 60}
	paco := Persona{"paco", 25}
	luis := Persona{"luis", 43}

	//Llamamos la funcion  Mayor
	tbMayor, tbDiff := Mayor(hugo, paco)
	tpMayor, tpDiff := Mayor(hugo, luis)
	bpMayor, bpDiff := Mayor(paco, luis)

	//Imprimimos los resultados
	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n",
		hugo.Nombre, paco.Nombre, tbMayor.Nombre, tbDiff)

	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n",
		hugo.Nombre, luis.Nombre, tpMayor.Nombre, tpDiff)

	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n",
		paco.Nombre, luis.Nombre, bpMayor.Nombre, bpDiff)
}