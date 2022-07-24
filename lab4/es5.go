/* Scrivere un programma trova.go che legge da standard input un char (rune)
e una stringa e stampa la posizione del char nella stringa,
o -1 se il char non c'è */

package main

import (
	"fmt"
	"strings"
)

func main () {

	var s string
	var b byte
	fmt.Scan(&s)
	fmt.Scanf("%c",&b)
	res := strings.Index(s,string(b))
	fmt.Print(res)
		
	
}