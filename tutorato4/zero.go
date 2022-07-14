/* Specifiche: Scrivere un programma massimo_n_zeri.go che legge da standard input una sequenza di interi
terminata da -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi siano più numeri
che contengono il massimo numero di 0, il programma stampa l’ultimo incontrato. Ad esempio, se la sequenza
letta è 3040 145 80 1707 105002 78 1970 6005 -1 il programma stampa 105002.
Per quanto riguarda l’input, osservate che: i dati arrivano da input standard; si tratta di una serie di dati
che possono essere elaborati in modo uniforme; non è necessario memorizzare l’intera sequenza; la lettura si
interrompe quando viene inserito -1. */

package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main () {
var num int
var init int
var counter int
fmt.Scan(&num)
init = num
j:= strconv.Itoa(init)
counter = strings.Count(j,"0")

for {
	fmt.Scan(&num)
	if num == -1 {
		break
	}
	x := strconv.Itoa(num)
	a:= strings.Count(x,"0")
	if a >= counter {
		counter = a
		fmt.Println(strconv.Atoi(x))
	}
	}
	
}
