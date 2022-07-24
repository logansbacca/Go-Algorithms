/* Scrivere un programma fibonacci.go che legge un intero
positivo n e stampa i numeri di fibonacci dal primo all'n-esimo,
rappresentandoli come righe di asterischi, ciascuna lunga quanto
il numero da rappresentare.
Esempio di esecuzione
un numero: 6
*
*
**
***
*****
********
------------------- */

// 1, 1 , 2, 3

package main

import (
	"fmt"
	"strings"
)

func main() {
	var in int
	var prev = 0
	var current = 1
	var x = 1
	fmt.Scan(&in)

	for i := 0; i<in; i++ {
		fmt.Println(strings.Repeat("*", x))
		x = prev + current
		prev = current
		current = x
	}

}
