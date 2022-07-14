package main

import (
	"fmt"
	"os"
	"strconv"
) 

func main() {
	var s [10]int
 	input := os.Args[1]
	fmt.Println(" Inserisci 10 numeri :")
	for i := 0; i < 10; i++ {
		fmt.Scan(&s[i])
	}
	n, _ := strconv.Atoi(input)
	p := -1
	for i, z := range s {
		if z == n {
			p = i
		}
	}
	fmt.Println(p)	
}

