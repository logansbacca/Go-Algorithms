/* Scrivere un programma minu_maiu.go che legge da standard input una stringa
e stabilisce se la stringa contiene solo minuscole, solo maiuscole o sia minuscole
che maiuscole, quindi stampa un messaggio opportuno (es. "contiene solo
minuscole", */

package main

import "fmt"

func main() {
	var str string
	var min int
	var max int
	fmt.Scan(&str)

	for _,v := range str {
		if v >= 'A' && v<='Z' {
			max++
		} else if v >= 'a' && v<='z' {
			min++
		}
	}

	if min > 0  && max == 0{
		fmt.Print("contiene solo minuscole")
	} else if max > 0 && min == 0{
		fmt.Print("contiene solo maiuscole")
	} else {
		fmt.Print("contiene sia min che maiusc")
	}

}