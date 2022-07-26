/* Scrivere un programma cesare.go che legge da standard input
un valore intero non negativo k (la chiave di cifratura) e
una sequenza di lettere minuscole consecutive
(sulla stessa riga e senza spazi) terminate da <invio> ('\n').
Il programma stampa la sequenza letta cifrata secondo il
cifrario di Cesare, usando come chiave k (quella fornita dall'utente):
ogni lettera del testo in chiaro è sostituita nel
testo cifrato dalla lettera che si trova k posizioni dopo
nell'alfabeto, ritornando dopo la zeta alla lettera a.
Esempi di esecuzione
chiave: 2
caratteri da cifrare: zaprb
bcrtd è il testo cifrato
chiave: 100
caratteri da cifrare: abcd
wxyz è il testo cifrato */

package main

import "fmt"

func main() {
	var letter rune
	var posizione int
	var k int
	var cifrato []rune
	alfabeto := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("inserire chiave")
	fmt.Scan(&k)
//zaprb
	for {
		fmt.Scanf("%c", &letter)
		if letter == '\n' {
			break
		}
		for i, v := range alfabeto {
			if v == letter {
				posizione = i
			}
		}
		nuovaposizione := ((posizione + int(k))% len(alfabeto)) 
		if nuovaposizione < posizione {
			cifrato = append(cifrato, rune(alfabeto[nuovaposizione]))
		}else {
			cifrato = append(cifrato, rune(alfabeto[nuovaposizione]))
		}
	}
	fmt.Print(string(cifrato))
}
