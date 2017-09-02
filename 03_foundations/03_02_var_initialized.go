package main

import "fmt"

func main() {
	var i, a, b int
	var j, k, l, m = 1, 2, "prueba", 1.2
	A, B, C := 'A', 8, 3.1416
	z := 6
	n := 9
	z, n = n, z
	fmt.Println("z:", z)
	fmt.Println("n:", n)

	_, gabo := "x", "gabo"
	fmt.Println(gabo)

	fmt.Println("lo que sea", i, a, b)
	fmt.Println("otra cosa:", j, k, l, m)
	fmt.Println("declaracion y asignacion:", A, B, C)
}
