package client_two

import "github.com/nicobianchetti/PatronesGo/Creacionales/01-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
