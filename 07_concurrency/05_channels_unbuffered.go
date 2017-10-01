package main

import (
	"fmt"
	"time"
	"strconv"
)

func recuperaMensaje (channel chan string) {
	var contador int
	for {
		contador++
		fmt.Println(<-channel,":", "recibiendo del channel",strconv.Itoa(contador))
		time.Sleep(time.Second * 1)
	}
}



func enviarMensaje (channel chan string) {
	var contador = 1
	for {
		channel <- "Mensaje enviado a Channel" + strconv.Itoa(contador)
		contador++
	}
}


func main() {

	// este channel solo recibe un valor
	ch := make(chan string)

	go enviarMensaje(ch)
	go recuperaMensaje(ch)

	var entrada_teclado string
	fmt.Scanln(&entrada_teclado)
	fmt.Print("Fin de ejecución")

}