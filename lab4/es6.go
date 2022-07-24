/* Scrivere un programma crescente.go che legge da standard input una stringa
di sole lettere minuscole e la stampa inserendo un'-' ogni volta
che una lettera viene prima in ordine alfabetico della lettera
precedente.
Per esempio, data in input la parola
ambire
il programma stampa
am-bir-e
================================================= */

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	var prec = str[0]
	var newarr [] string
	for i:=0; i<len(str);i++{
		newarr = append(newarr, string(str[i]))
		if prec < str[i]{
			prec= str[i]
			newarr = append(newarr, "-")
		}
	}

	for _,v := range newarr {
		fmt.Print(v)
	}
	
}