package main

import "fmt"


// PARA DEMOSTRAR LA INCORPORACION DE UN METODO A UN TIPO
// Y COMO EL RECEPTOR PUEDE COMPARARSE CON EL this EN JAVA
type dosEnteros struct {
	a int
	b int
}

func (this *dosEnteros) sumaNumeros() int {
	return this.a + this.b
}

func main() {
	var i = new(dosEnteros)
	i.a = 3
	i.b = 7
	fmt.Println(i.sumaNumeros())
}
