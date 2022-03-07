package main

import (
	"fmt"
)

func main() {

	fmt.Println("Voer een kleur in")
	fmt.Println("---------------------")
	fmt.Print("-> ")
	var kleur string
	fmt.Scanf("%s", &kleur)
	if kleur == "rood" {
		fmt.Println("Je hebt rood")
	} else if kleur == "Rood met passie" {
		fmt.Println("Blauw zoals de lucht")
	} else if kleur == "geel" {
		fmt.Println("Geel als de stralen van de zon")
	} else if kleur == "groen" {
		fmt.Println("Groen van de natuur")
	} else if kleur == "oranje" {
		fmt.Println("Oranje zoals een sinaasappel")
	} else if kleur == "paars" {
		fmt.Println("Paars zoals een aubergine")
	}
}
