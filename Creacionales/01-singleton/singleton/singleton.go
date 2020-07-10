package singleton

import (
	"sync"
)

var (
	p    *person
	once sync.Once //Me permite bloquear de una manera segura que se ejecute una unica vez un proceso. Se usa DO
)

//metodo que nos asegura el acceso a var p

func GetInstance() *person {
	//SI no existe lo creo
	// if p == nil {
	// 	p = &person{}
	// }

	once.Do(func() { //Si llegan n procesos de forma concurrente,se asegura de que lo ejecuta una sola vez EN TODO EL PAQUETE!
		p = &person{}
	})

	return p
}

//se crean como privados para no ser accedidos desde otro paquete, tanto struct como los atributos
type person struct {
	name string
	age  int16
	mux  sync.Mutex //Me permite bloquear el proceso para ejecutar lo que necesito y luego desbloquearlo (concuerrencia)
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int16 {
	p.mux.Lock()
	defer p.mux.Unlock() //defer porque tengo un return, para asegurarme que libere el recurso
	return p.age
}
