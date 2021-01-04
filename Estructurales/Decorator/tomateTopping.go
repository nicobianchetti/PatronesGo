package main

type tomateTopping struct {
	pizzaTomate pizza
}

func (t *tomateTopping) getPrice() int {
	pizzaPrice := t.pizzaTomate.getPrice()
	return pizzaPrice + 7
}
