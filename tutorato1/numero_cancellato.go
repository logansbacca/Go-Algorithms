
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers string
	var toint []int
	var intarr int
	fmt.Scan(&numbers)
	chars := []rune(numbers)
	toint = make([]int, 0)
	var try []  rune
	try = append(try,chars[len(chars)-1])


	//trasformo gli asterischi in 0 e il numero prima del asterisco anche
	for i := 0; i < len(chars)-1; i++ {
		if chars[i+1] == 35 || chars[i] == 35 {
			continue
		} else {
			try = append(try,chars[i])
		}
		
	}


// converto i numeri di stringhe in int e li aggiungo ad un nuovo array di interi
	for j := 0; j < len(try); j++ {
		intarr, _ = strconv.Atoi(string(try[j]))
		toint = append(toint, intarr)
	}


	// itero per multiplicare
	product := 1
	for i := 0; i < len(toint); i++ {
		product = product * toint[i]
	}
	fmt.Println(product,toint)
}
