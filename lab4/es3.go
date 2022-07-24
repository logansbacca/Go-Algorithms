/* Scrivere un programma lunghezza_tot.go che legge da standard input
un int totLen e una sequenza di stringhe (una per riga) e
ne somma le lunghezze fino a raggiungere (o superare) totLen.
Quanto termina, stampa la somma delle lunghezze delle stringhe lette.
*/

package main 

import "fmt"

 func main () {
	var stringsum int
	var totlen int
	var current string
	fmt.Scan(&totlen)
	stringsum = len(current)
	
	for{
		if stringsum >= totlen {
			break
		}
		fmt.Scan(&current)
		stringsum = stringsum + len(current)
	}

	fmt.Print(stringsum,totlen)

 }