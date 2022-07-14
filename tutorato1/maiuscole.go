package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	const BASE_MINU = 'a'
	const BASE_MAIU = 'A'
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserisci la stringa da mettere in maiuscolo :")
	if scanner.Scan() {
		s = scanner.Text()
	}
	fmt.Println("La stringa è lunga", len(s))
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			n := char - BASE_MINU
			char = BASE_MAIU + n
		}
		fmt.Print(string(char))
	}
	fmt.Println()
}

/*1. Il programma maiuscole.go attraverso quale canale riceve i dati da elaborare? 
2. Di che tipo sono i dati ricevuti in input dal programma? 
di tipo stringa
3. I dati ricevuti da input sono pronti da elaborare o vengono manipolati prima in qualche modo (per calcolarne/derivarne/estrarne altri dati)?
Il for loop serve per iterare su ogni carattere e viene controllata la condizione secondo la quale poi al carattere viene sottratto il numero di "a"  ed assegnato ad n. Poi char assume il valore di "A" piu n. Dato il valore sara in rune viene poi convertito a string.
\*/
