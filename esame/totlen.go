package main

/* Scrivere un programma lunghezza_tot.go che legge da standard input
un int totLen e una sequenza di stringhe (una per riga) e
ne somma le lunghezze fino a raggiungere (o superare) totLen.
Quanto termina, stampa la somma delle lunghezze delle stringhe lette */

import "fmt"

func main () {
	var parola string
	totlen:= 10
	var parole [] string
	for {
		fmt.Scan(&parola)
		parole = append(parole,parola)
		if len(parole) == totlen {
			break
		}
	}
	
}