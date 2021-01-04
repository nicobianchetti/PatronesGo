package main

import "fmt"

type windows struct{}

func (w *windows) inserIntoUSBPort() {
	fmt.Println("USB connector is plugged into Windows machine")
}
