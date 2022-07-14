/* Scrivere un programma num2text.go per convertire un numero
intero non negativo nella sequenza delle parole corrispondenti
alle sue cifre.
Il programma legge un intero non negativo da standard input,
per ogni nuova (non incontrata finora) cifra del numero
chiede il nome corrispondente (e alimenta un dizionario), e infine
stampa la sequenza delle parole corrispondenti alle sue cifre.
Ad esempio, per il numero 203, il programma stampa
due - zero - tre

Esempio di esecuzione
---------------------
$ go run num2text.go
un numero: 622
parola per 2 ? due
parola per 6 ? sei
sei - due - due */

package main

import (
	"fmt"
)

func main() {
	var num int
	var x int
	var res []int

	dizionario := map[int]string{
		0: "zero",
		1: "uno",
		2: "due",
		3: "tre",
		4: "quattro",
		5: "cinque",
		6: "sei",
		7: "sette",
		8: "otto",
		9: "nove",
	}

	fmt.Scan(&num)

	for num > 0 {
		x = num % 10
		num = num / 10
		res = append(res, x)
	}
	

 	var rev [] int
	for i:=len(res)-1; i>=0; i-- {
		rev = append(rev,res[i])
	}
	
	for i:=0; i<len(rev); i++ {
		fmt.Print(dizionario[rev[i]])
		if i != len(rev)-1 {
			fmt.Print(" - ")
		}
	}
}