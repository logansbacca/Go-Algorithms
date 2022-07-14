/* Specifiche: Scrivere un programma che legga una sequenza di N interi (almeno due) e stampi la più grande
differenza (in valore assoluto) che c’è tra due numeri consecutivi. */

package main

import "fmt"

func main() {
	var num int
	var first int
	var res int
	var dif int
	fmt.Scan(&num)
	first = num
	for {
		fmt.Scan(&num)
		dif = first - num
		if dif < 0 {
			break
		}
		if dif> res {
			res = dif
		}
		first = num
	}
	fmt.Println(res)
}
