package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Go Factory")

	ak47, _ := GetFactory().getOption(EnumOptionAk47)

	// ak47, _ := getGun("ak47")
	// m4, _ := getGun("m4")

	printDetails(ak47)
	// printDetails(m4)

}

func printDetails(gun iGun) {
	fmt.Printf("Gun: %s", gun.getName())
	fmt.Println()
	fmt.Printf("Power: %d", gun.getPower())
	fmt.Println()
}
