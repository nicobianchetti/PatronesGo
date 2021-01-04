package main

type quesoTopping struct {
	pizzaQueso pizza
}

func (q *quesoTopping) getPrice() int {
	pizzaPrice := q.pizzaQueso.getPrice()
	return pizzaPrice + 10
}
