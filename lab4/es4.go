/* Scrivere un programma max_char.go che legge da standard input
una sequenza di 5 caratteri ASCII (byte) e stampa il maggiore
in ordine lessicografico (codice ASCII)./*  */

package main

import (
	"fmt"
	"sort"
)

func main () {
	var input byte
	var order [] string

	for i:=1; i<=5; i++ {
		fmt.Scanf("%c",&input)
		order = append(order, string(input))
	}

	sort.Strings(order)
	fmt.Print(order)

}