package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	fmt.Println(numVocali(input))
}

func numVocali(parole []string) string {
	init := 0
	var counter int
	var res string

	for i := 0; i < len(parole); i++ {
		counter = 0
		parola := parole[i]
		for j := 0; j < len(parola); j++ {
			if parola[j] == 'a' || parola[j] == 'e' || parola[j] == 'i' || parola[j] == 'o' || parola[j] == 'u' {
				counter++
			}
		}
		if counter > init {
			init = counter
			res = parola
		}

	}
	return res
}
