/* Scrivere un programma conta_vocali.go che legge da standard input
una parola e stampa il numero di lettere (rune) che sono vocali (aeiou) */

package main

import "fmt"

func main () {
	var parola string
	vocali :="aeiou"
	var counter int
	fmt.Scan(&parola)
	
	for _,v := range vocali {
		for _,b  := range parola {
			if v == b {
				counter++
			}
		}
	}
	fmt.Print(counter)
}