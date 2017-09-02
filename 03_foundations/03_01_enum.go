package main

import "fmt"

// Constantes de cursos PaaS Studio
const (
	Ansible = iota
	Docker
	Go
)

func main() {
	const x string = "Hola Mundo desde una constante"
	fmt.Println(x)
	fmt.Println("Ansible:", Ansible, "\nDocker: ", Docker, "\nGo:", Go)

}
