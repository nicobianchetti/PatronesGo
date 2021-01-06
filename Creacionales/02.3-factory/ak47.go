package main

import (
	"fmt"
	"log"
)

const EnumOptionAk47 = "ak47"

type ak47 struct {
	gun
}

func init() {
	addOption(&ak47{})
	fmt.Println(EnumOptionAk47 + " added to factory")
}

func (a *ak47) validate(params ...interface{}) bool {
	log.Println("params Ak47")
	log.Println(params...)
	if params[0] == EnumOptionAk47 {
		return true
	}
	return false
}

func (a *ak47) getOption() IFactoryOption {
	return &ak47{}
}

func (a *ak47) Do() {
	fmt.Println("Do: Option ak47")
}

// func newAk47() iGun {
// 	return &ak47{
// 		gun: gun{
// 			name:  "AK47 gun",
// 			power: 4,
// 		},
// 	}
// }
