package main

type iGun interface {
	setName(name string)
	setPower(power int)
	getPower() int
	getName() string
}
