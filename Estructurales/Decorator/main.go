package main

import "fmt"

func main() {
	pizza := &napolitana{}

	//Add queso topping
	pizzaConQueso := &quesoTopping{
		pizzaQueso: pizza,
	}

	//Add tomate topping
	pizzaConTomate := &tomateTopping{
		pizzaTomate: pizzaConQueso,
	}

	fmt.Printf("El precio de la pizza Napolitana con queso y tomate es %d\n", pizzaConTomate.getPrice())
}
