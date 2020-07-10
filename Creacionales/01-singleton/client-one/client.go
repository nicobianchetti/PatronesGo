package client_one

import "github.com/nicobianchetti/PatronesGo/Creacionales/01-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
