package main

import "fmt"

func main() {
	var s [10]int
	var n int

	fmt.Println(" Inserisci 10 numeri :")
	for i := 0; i < 10; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println(" Inserisci un numero :")
	fmt.Scan(&n)

	var i int
	for i = 9; i >= 0; i-- {
		if s[i] == n {
			break
		}
	}
	fmt.Println(i)
}


/* Analizzate il codice sorgente e rispondete per iscritto alle domande che seguono.
1. Il programma array.go attraverso quale canale riceve i dati da elaborare?
riceve i dati via scan
2. Di che tipo/i sono i dati in input?
sono numeri interi
3. I dati ricevuti da input sono pronti da elaborare o vengono manipolati prima in qualche modo (per
calcolarne/derivarne/estrarne altri dati)?
dall array s viene analizzato ogni numero nella sequenza
4. Quali, fra i dati in input, sono memorizzati in variabili singole e quali in strutture (array, slice, ecc.)?
l'indice viene salvato nella varabile singola e i 10 numeri nell array s.
5. Quando si ferma il programma nella lettura della sequenza di numeri?
quando trova il numero nel array s che corrisponde ad n
6. Perché è necessario salvare alcuni dati in una struttura?
per poterli confrontare con n
\*/