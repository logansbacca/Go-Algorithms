/* Scrivere un programma ultima_pioggia.go che legge da standard input una
sequenza di valori interi (terminata da EOF, da tastiera prodotto con ctrl-D)
che indicano i mm di pioggia caduti (0 se non ha piovuto) ogni giorno
in una sequenza successiva di giorni e stampa (l'indice del) l'ultimo
giorno in cui ha piovuto.
*/

package main

import "fmt"

func main () {
	var mm int
	var ultimo int
	for {
		_,err := fmt.Scan(&mm)
		if err != nil {
			break
		}
		if mm > 0 {
			ultimo = mm
		}
	 }
	 fmt.Print(ultimo)
}



