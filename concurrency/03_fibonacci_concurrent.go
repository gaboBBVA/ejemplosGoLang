package main

import (
	"time"
	"fmt"
)

func main() {
	go procesando(100*time.Millisecond)
	calculo_fibonacci := fibonacci(400)
	fmt.Println("\nEl fibonacci calculado es:",calculo_fibonacci)
}

func procesando (duracion time.Duration) {
	for {
		for _, caracter := range `\|/-` {
			fmt.Printf("\r %c", caracter)
			time.Sleep(duracion)
		}
	}
}


func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}