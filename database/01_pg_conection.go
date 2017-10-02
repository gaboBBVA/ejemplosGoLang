package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "golangdb"
)


// TABLA DE USUARIOS EN LA BASE DE DATOS POSTGRESQL
type Usuario struct {
	ID          int
	Nombre       string
	Password string
	Correo       string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM prueba")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var usuarios []*Usuario
	for rows.Next() {
		usr := &Usuario{}
		err := rows.Scan(&usr.ID, &usr.Nombre, &usr.Password, &usr.Correo)
		if err != nil {
			log.Fatal(err)
		}
		usuarios = append(usuarios, usr)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, pr := range usuarios {
		fmt.Printf("%d, %s, %s, %s\n", pr.ID, pr.Nombre, pr.Password, pr.Correo)
	}

}
