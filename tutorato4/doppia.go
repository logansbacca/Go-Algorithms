/* Scrivere un programma che riceve da riga di comando una serie di lunghezza arbitraria di parole
e stampa la prima parola che contiene una doppia (una lettera ripetuta due volte consecutivamente). Nel caso
nessuna parola contenga una doppia, il programma deve stampare un messaggio opportuno. */

package main

import "fmt"
import "os"

func main () {
	var lettera string
	input := os.Args[1:]
	var message bool
	for _,v  := range input {
		lettera = v
		for i:=0; i<len(lettera)-1; i++{
			if lettera[i] == lettera[i+1]{
				message = true
				fmt.Print(lettera)
				break
			}
		}
		}
		if !message {
			fmt.Print("non ci sono doppie ",message)
		}
	}
