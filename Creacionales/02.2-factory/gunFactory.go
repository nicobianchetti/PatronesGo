package main

import (
	"errors"
)

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "m4":
		return newM4(), nil
	default:
		err := errors.New("Not found Gun")
		return nil, err
	}
}
