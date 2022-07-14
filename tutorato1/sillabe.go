package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(dividiInSillabe(os.Args[1]))
}
func dividiInSillabe(parola string) (sliceSillabe []string) {
	var letteraPrecedente rune
	var sillaba string
	for _, letteraCorrente := range parola {
		if letteraCorrente < letteraPrecedente {
			sliceSillabe = append(sliceSillabe, sillaba)
			sillaba = string(letteraCorrente)
		} else {
			sillaba += string(letteraCorrente)
		}
		letteraPrecedente = letteraCorrente
	}
	sliceSillabe = append(sliceSillabe, sillaba)
	return
}
