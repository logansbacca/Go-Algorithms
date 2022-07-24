/* Scrivere un programma es0.go che
- legge un byte e
- lo stampa (occorre Scanf in più per catturare l'invio)
- stampa il byte precedente, il byte stesso, e quello successivo
in ordine lessicografico (ASCII). Ad es. per 'd', deve stampare: cde
- stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
- poi legge una stringa e la stampa in verticale, una runa per riga. Ad esempio città:
c
i
t
t
à */

package main

import (
	"fmt"

)

func main () {
	var input byte 
	fmt.Scanf("%c",&input)
	prev := input-1
	next := input +1
	fmt.Println(prev,input, next)

	if input >= 'a' && input <= 'l'{
		fmt.Println("a-l")
	} else {
		fmt.Print("altro")
	} 
	var word string
	fmt.Scan(&word)
	for _,v := range word {
		fmt.Println(string(v))
	}
}