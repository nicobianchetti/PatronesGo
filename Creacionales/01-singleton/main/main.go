package main

import (
	"fmt"
	"sync"

	client_one "github.com/nicobianchetti/PatronesGo/Creacionales/01-singleton/client-one"
	client_two "github.com/nicobianchetti/PatronesGo/Creacionales/01-singleton/client-two"
	"github.com/nicobianchetti/PatronesGo/Creacionales/01-singleton/singleton"
)

func main() {
	/*
		//Este ejemplo es sincrónico, no hay concurrencia.
		client_one.IncrementAge()
		client_two.IncrementAge()
		p := singleton.GetInstance()
		age := p.GetAge()
		fmt.Println(age)
	*/

	//Ejemplo concuerrencia. Uso paquete wait group
	wg := sync.WaitGroup{}
	wg.Add(200) //200 ejecuciones

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}

	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
