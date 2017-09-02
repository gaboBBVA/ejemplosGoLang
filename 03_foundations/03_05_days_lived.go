package main

import "fmt"

func main() {

	var (
		nombre string
		años   int
	)
	//var nombre string
	fmt.Print("¿Cuál es tu nombre?:")
	fmt.Scanf("%s", &nombre)

	//var años int
	fmt.Print("¿Cuántos años tienes?:")
	fmt.Scanf("%d", &años)

	fmt.Println("Hola ", nombre, " tienes por lo menos ", años*365, "días vivídos")
}
