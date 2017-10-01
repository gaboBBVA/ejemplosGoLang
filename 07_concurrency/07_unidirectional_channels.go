package main

import (
	"fmt"
	"time"
)

func genera_numeros (write_channel_genera_numero chan int) {
	for x := 0; x < 5; x++ {
		write_channel_genera_numero <- x
	}
	close (write_channel_genera_numero)
}

func elevar_al_cuadrado (read_channel_genera_numero chan int, write_channel_cuadrado chan int) {
	for x := range read_channel_genera_numero {
		write_channel_cuadrado <- x * x
		// TODO cambiar el sentido de un channel compila pero el error se genera en tiempo de ejecucion
		// read_channel_genera_numero <- 5
		// TODO cambiar la declaracion de las funciones para que sean unidireccionales y evitar error de ejecucion
		// write chan<-
		// read <-chan
	}
	close(write_channel_cuadrado)
}

func imprimir_resultados(read_channel_cuadrado chan int) {
	var i int
	for x := range read_channel_cuadrado {
		fmt.Println("el cuadrado de: ",i,"es:",x)
		time.Sleep(time.Second * 1)
		i++
	}
}

func main (){

	channel_genera_numeros := make(chan int)
	channel_eleva_al_cuadrado := make(chan int)

	go genera_numeros(channel_genera_numeros)
	go elevar_al_cuadrado(channel_genera_numeros,channel_eleva_al_cuadrado)

	imprimir_resultados(channel_eleva_al_cuadrado)

}