package main

import (
	"net"
	"log"
	"io"
	"time"
)

func managerConexiones (conexion net.Conn) {

	for {
		_, err := io.WriteString(conexion, time.Now().Format(time.RFC3339) + "\n\r")

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	defer conexion.Close()
}


func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conexion, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// managerConexiones(conexion)
		go managerConexiones(conexion)
	}
}