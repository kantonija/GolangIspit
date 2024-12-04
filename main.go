package main

import "fmt"

func main() {
	proizvodi := map[string]int{
		"jabuka":     10,
		"mlijeko":    50,
		"svinjetina": 170,
		"janjetina":  210,
		"banana":     30,
	}

	Proizvodi(proizvodi)
}

func Proizvodi(proizvodi map[string]int) {
	var max int = 0
	var maxIme string
	// var min int = 0
	for ime, cijena := range proizvodi {
		if cijena > max {
			max = cijena
			maxIme = ime
		}
	}
	fmt.Printf("Najskuplji proizvod je %s s cijenom od %d.\n", maxIme, max)

	fmt.Println("Proizvodi cijena koja je veca od 150:")
	for ime, cijena := range proizvodi {
		if cijena > 150 {
			fmt.Printf("%s - %d\n", ime, cijena)
		}
	}

}
