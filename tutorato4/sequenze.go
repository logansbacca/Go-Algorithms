/* Specifiche: Scrivere un programma che legge una sequenza di interi e stampa la somma degli interi in ciascuna
sottosequenza crescente. Il programma interrompe la lettura quando riceve un segnale di EOF. Ad esempio,
su input 1 2 13 0 7 8 9 -1 0 2 il programma stampa le somme 16, 24 e 1. */

package main

import (
	"fmt"
	"io"
)

func main() {
	var input int 
	var sum int
	numarr:= [] int {}

	for {
		_,err := fmt.Scan(&input)
		if err == io.EOF{
			break
		}
		numarr = append(numarr, input)
	}

	for i:=0; i<len(numarr)-1; i++ {
		if numarr[i]<numarr[i+1]{
			sum = sum + numarr[i]
		} else if numarr[i] > numarr [i+1]{
			sum= sum + numarr[i]
			fmt.Println(sum)
			sum= 0
			sum = sum + numarr[i]
			sum = 0
		}
	}
	fmt.Println(sum)

}