package main

import (
	"time"
	"os"
	"fmt"
)

func leer_entrada_teclado (write_channel_teclado chan<- string) {
	for {
		data := make([]byte, 1024)
		n, _ := os.Stdin.Read(data)
		var s string
		s = string(data[:n])
		if n > 0 {
			write_channel_teclado <- s
		}
	}
}

func main() {
	fin := time.After(time.Second * 30)
	eco := make (chan string)
	go leer_entrada_teclado(eco)

	for {
		select {
			case datos := <-eco:
				//os.Stdout.Write(datos)
				fmt.Printf("%s", datos)
			case <-fin:
				fmt.Println("\nSe terminó el tiempo")
				os.Exit(0)
		}
	}
}