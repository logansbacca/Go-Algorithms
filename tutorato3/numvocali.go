package main

import "fmt"

func main () {
	fmt.Println(numVocali("albero"))
}

func numVocali (parola string) int {
	var counter int
	for i:=0; i<len(parola); i++ {
			if string(parola[i]) == "a" || string(parola[i]) == "e" || string(parola[i])== "i" || string(parola[i]) == "o" || string(parola[i]) == "u" {
				counter++
				continue
			}
		}
	return counter
}