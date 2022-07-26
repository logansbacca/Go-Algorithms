/* Scrivere un programma num_sequenze.go che legge da standard input una
sequenza di uni (1) e zeri (0) (terminata da un 2), che inizia e finisce con 1,
e stampa il numero di sottosequenze di zeri.
Ad esempio per input 1 1 0 0 1 0 1 0 0 0 1 1 1 0 1,
l'output è 4.
(si considerano anche quelle di lunghezza 1) */

package main

import "fmt"

func main () {
	var sub int
	var subs [] int
	var count int
	for{
		
		_,err := fmt.Scan(&sub)
		if err != nil {
			break
		}
		subs = append(subs, sub)
	}

	for i:=0; i<len(subs); i++{
		if subs[i] == 0 && subs [i+1] == 1 {
			count++
		}
	}
	fmt.Print(count)
}