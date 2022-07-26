/* Scrivere un programma max_num_cifre_pari.go che data una sequenza di numeri
(letti come stringhe),
stampa il massimo numero di cifre pari contenute in un numero */

package main

import (
	"fmt"
	//"strconv"
)

func main () {
	var maxnum = []string {"733", "667", "12", "444"}
	var counter int

	for i,v := range maxnum {
		for _,x := range v {
			if x % 2 == 0 {
				counter++
			}
			/* a,_:= strconv.Atoi(string(x))
			if a % 2 == 0 {
				counter++
			} */
		}
		fmt.Println("word number", i+1, "has : ", counter, "amount of even nums" )
		counter = 0
		
	}

}