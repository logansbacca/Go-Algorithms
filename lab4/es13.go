/* Scrivere un programma num_max.go che legge una sequenza di 10 interi positivi
e stampa il massimo intero letto e quante volte il massimo compare nella sequenza. */

package main

import "fmt"

func main() {
	var in int
	var previn int
	var biggest int
	var sec []int
	var counter int
	fmt.Scan(&in)
	previn = in
	for i := 0; i < 10; i++ {
		fmt.Scan(&in)
		sec = append(sec, in)
		if previn < in {
			previn = in
		}
	}
	biggest = previn

	for _, v := range sec {
		if v == biggest {
			counter++
		}
	}
	fmt.Print(biggest, counter)
}
