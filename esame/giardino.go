/* Definiamo "gradino" una sequenza di (uno o più) interi non negativi uguali seguita
da un'altra sequenza di (uno o più) interi più grandi di 1 (es. 1 1 2 2 2).
Scrivere un programma gradino.go che, data in input una sequenza di interi
tali che ogni intero è >= del precedente, stampa la lunghezza (il numero di interi)
del gradino più lungo. (Si noti che i gradini si sovrappongono).
Il programma termina quando legge un numero minore di quello appena letto.
Per input 2 2 3 3 4 4 4 5 6 6 6 7, l'output è 5 (il gradino 3 3 4 4 4).
________________________________ */

package main

import "fmt"

func main (){
	var current, prev int
	var currlen int
	var prevlen int
	maxStep := 0
	fmt.Scan(&current)
	for current >= prev{
		prev = current
		fmt.Scan(&current)
		if current == prev {
			currlen ++
		} else  {
			currstep := prevlen + currlen
			if maxStep < currstep{
				maxStep = currstep
			} 
			prevlen = currlen
			currlen = 1
		}
	}
	fmt.Print(maxStep)
}