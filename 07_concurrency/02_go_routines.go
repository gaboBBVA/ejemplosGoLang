package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var espera_grupo sync.WaitGroup

func imprime_numeros(id string, iteraciones int) {
	for i:=1;i<=iteraciones;i++ {
		retardo:= rand.Int63n(1000)
		time.Sleep(time.Duration(retardo) * time.Millisecond)
		fmt.Printf("%s -> %d\n",id,i)
	}
	defer espera_grupo.Done()
}

func main() {
	espera_grupo.Add(4)

	fmt.Println("Ejecución de la primer go-routine : A,10")
	go imprime_numeros("A",10)

	fmt.Println("Ejecución de la segunda go-routine : B,10")
	go imprime_numeros("B",10)

	fmt.Println("Ejecución de la segunda go-routine : C,10")
	go imprime_numeros("C",10)

	fmt.Println("Ejecución de la segunda go-routine : D,10")
	go imprime_numeros("D",10)

	fmt.Println("Esperando a que finalicen las go-routines")
	espera_grupo.Wait()

	fmt.Println("Fin de la ejecución")

}