package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

var (
	KleineLetterSet = "abcdedfghijklmnopqrst"
	HoofdLetterSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecKarakterSet = "!@#$%&*"
	NummerSet       = "0123456789"
	allCharSet      = KleineLetterSet + HoofdLetterSet + SpecKarakterSet + NummerSet
)

func main() {
	Titel := figure.NewColorFigure("Wachtwoord Generator", "", "green", true)
	Titel.Print()
	Naam := figure.NewColorFigure("Gemaakt door Melvin Thijssen", "mini", "red", true)
	Naam.Print()
	rand.Seed(time.Now().Unix())
	var speckar = 0
	var Nummers = 0
	var hoofdletters = 0
	var Lengte = 0

	fmt.Println("Hoeveel nummers?: ")
	fmt.Scanln(&Nummers)
	fmt.Println("Hoeveel speciale karakters?: ")
	fmt.Scanln(&speckar)
	fmt.Println("Hoeveel hoofdletters?: ")
	fmt.Scanln(&hoofdletters)
	fmt.Println("Hoe lang?: ")
	fmt.Scanln(&Lengte)

	fmt.Println("----- Wachtwoord 1 ------")
	password := Genereer(Lengte, speckar, Nummers, hoofdletters)
	fmt.Println(password)
	fmt.Println("-------------------------")

	fmt.Println("----- Wachtwoord 2 ------")
	password = Genereer(Lengte, speckar, Nummers, hoofdletters)
	fmt.Println(password)
	fmt.Println("-------------------------")
}

func Genereer(Lengte, speckar, Nummers, hoofdletters int) string {
	var wachtwoord strings.Builder

	for i := 0; i < speckar; i++ {
		random := rand.Intn(len(SpecKarakterSet))
		wachtwoord.WriteString(string(SpecKarakterSet[random]))
	}

	for i := 0; i < Nummers; i++ {
		random := rand.Intn(len(NummerSet))
		wachtwoord.WriteString(string(NummerSet[random]))
	}

	for i := 0; i < hoofdletters; i++ {
		random := rand.Intn(len(HoofdLetterSet))
		wachtwoord.WriteString(string(HoofdLetterSet[random]))
	}

	remainingLength := Lengte - speckar - Nummers - hoofdletters
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		wachtwoord.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(wachtwoord.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
