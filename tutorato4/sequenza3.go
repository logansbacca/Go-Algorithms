//Scrivere un programma che riceve da riga di comando una serie di lunghezza arbitraria di
//parole e stampa la parola più lunga e la parola più corta. Nel caso in cui vi siano più parole di lunghezza
//massima o minima, il programma stampa l’ultima incontrata.

package main 

import "os"
import "fmt"

func main () {
	input := os.Args[1:]
	var lunga = input[1]
	var corta = input[1]
	for i:=0; i<len(input); i++ {
		if len(input[i]) >= len(lunga) {
			lunga = input[i]
		} else if len(input[i]) <= len(corta)  {
			corta = input[i]
		} 
		
	}

	fmt.Println(corta, lunga)
}