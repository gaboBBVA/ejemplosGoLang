package main

import "fmt"

func imprime_numeros(id string, iteraciones int) {
	for i:=1;i<=iteraciones;i++ {
		fmt.Printf("%s -> %d\n",id,i)
	}
}

func main() {
	imprime_numeros("caso A", 10)
	imprime_numeros("caso B", 10)
}


