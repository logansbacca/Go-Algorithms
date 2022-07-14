package main

import "fmt"

func main() {
	var fahr float64
	const FATTORE_SCALA, ZERO = 5.0 / 9.0, 32.0

	fmt.Println(" Inserisci la temperatura in gradi fahrenheit : ")
	fmt.Scan(&fahr)
	cels := (fahr - ZERO) * FATTORE_SCALA
	fmt.Println(" Gradi Celsius equivalenti :", cels)
}
