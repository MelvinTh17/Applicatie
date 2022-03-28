package main

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	Titel := figure.NewColorFigure("Palindroom checker", "", "green", true)
	Titel.Print()
	Naam := figure.NewColorFigure("Gemaakt door Melvin Thijssen", "mini", "red", true)
	Naam.Print()
	var Invoer string = ""
	fmt.Println("Wat wil je controleren?: ")
	fmt.Scanln(&Invoer)
	var Omgedraaid string = ""
	var lengte = len(Invoer)

	for i := lengte - 1; i >= 0; i-- {
		Omgedraaid = Omgedraaid + string(Invoer[i])
	}

	if strings.ToLower(Invoer) == strings.ToLower(Omgedraaid) {

		wit := color.New(color.FgWhite)
		GroenAchter := wit.Add(color.BgGreen)
		GroenAchter.Println("Dit is een mooie palindroom!")

	} else {
		wit := color.New(color.FgWhite)
		RoodAchter := wit.Add(color.BgRed)
		RoodAchter.Println("Dit is geen palindroom!")
	}
}
