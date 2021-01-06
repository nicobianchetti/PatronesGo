package main

import (
	"fmt"
	"os"

	"github.com/nicobianchetti/PatronesGo/Creacionales/02-factory/factory"
)

func main() {
	var t int
	fmt.Println("Digite la conexión que desea:\n Opción 1:MySql \n Opción 2:Postgres")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer la opción: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor no válido")
		os.Exit(1)
	}

	err = conn.Connect()

	if err != nil {
		fmt.Printf("No se pudo crear conexión: %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()

	if err != nil {
		fmt.Printf("no se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))
	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexión: %v", err)
	}
}
