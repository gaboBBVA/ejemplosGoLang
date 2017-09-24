package main

import (
	"fmt"
	"time"
)

func main() {
	channel_numero := make(chan int)
	channel_cuadrado := make(chan int)

	go func() {
		for x := 1; ; x++ {
			channel_numero <- x
			fmt.Print("\nenviado a channel_numero:", x)
		}
		// TODO cerrar el canal y ver que pasa con el resto
		//close(channel_numero)
	}()

	go func() {
		for {
			x := <-channel_numero
			fmt.Println("\nrecuperando de channel_numero:", x)
			cuadrado := x * x
			channel_cuadrado <- cuadrado
			fmt.Println("\nenviado a channel_cuadrado:", cuadrado)
		}
	}()

	for {
		fmt.Println("\nrecibiendo de channel_cuadrado:",<-channel_cuadrado)
		time.Sleep(time.Second * 1)
	}

}