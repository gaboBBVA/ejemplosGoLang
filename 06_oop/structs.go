package main

import "fmt"

type estructuraX struct {
	i int
	j float32
}

func main() {
	var x estructuraX
	x.i = 2
	x.j = 2.3
	fmt.Println("Valores:", x.i, x.j)

	y := new(estructuraX)
	y.i = 10
	y.j = 4.4
	fmt.Println("Valores:", y.i, y.j)
}
