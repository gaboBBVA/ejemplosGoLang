package main

import "fmt"

func main() {
	var nombre, apellido string

	fmt.Print("¿Cuál es tu nombre y primer apellido?:")
	fmt.Scanln(&nombre, &apellido)
	//fmt.Scanf("%s %s", &nombre, &apellido)
	fmt.Println("Hola ", nombre, apellido)
}
