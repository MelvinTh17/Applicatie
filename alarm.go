package main

import (
	"fmt"
)

func main() {

	fmt.Println("Voer een kleur in")
	fmt.Println("---------------------")
	fmt.Print("-> ")
	var kleur string
	fmt.Scanf("%d", &kleur)
	if kleur == "rood" {
		fmt.Println("Je hebt rood")

	}
}
