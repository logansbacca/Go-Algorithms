package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	a := numeroDaIndovinare(1, 100)
	for {
		fmt.Scan(&input)
		if input>a {
			fmt.Println("+")
		}else if input<a{
			fmt.Println("-")
		}else {
			fmt.Println("Hai indovinato, il numero era: ",a)
			break
		}
	}

}

func numeroDaIndovinare(lower, upper int) int {
	RandomIntegerwithinRange := rand.Intn(upper-lower) + lower
	return RandomIntegerwithinRange
}
