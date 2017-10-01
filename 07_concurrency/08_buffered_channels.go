package main

import (
	"strconv"
	"fmt"
	"time"
)

func enviar_mensaje (write_channel_mensajes chan<- string, numero int) {
	for {
		write_channel_mensajes <- "Mensaje puesto por la goroutine:" + strconv.Itoa(numero)
		fmt.Println("-> Escrito en el channel por parte de la goroutine:" + strconv.Itoa(numero))
	}
}

func imprimir(read_channel_mensajes <-chan string) {
	for x := range read_channel_mensajes {
		fmt.Println("<- Mensaje recuperado:",x)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	channel := make(chan string, 5)
	for i := 1; i < 6; i++ {
		go enviar_mensaje(channel,i)
	}
	imprimir(channel)
}