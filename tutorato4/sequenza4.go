// Scrivere un programma che legge da standard input una sequenza di interi terminata da
// -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi siano più numeri che
// contengono il massimo numero di 0, il programma stampa l’ultimo incontrato. Ad esempio, se la
// sequenza letta è 3040 145 80 1707 105002 78 1970 6005 -1 il programma stampa 105002.

package main

import (
	"fmt"
	"strconv"
	
)

func main()  {
	var input int
	var n string
	var numbers []string
	var number string
	var zerocounter int
	var currentcounter int
	var longest string
	var longestnum int64


	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		n = strconv.Itoa(input)
		numbers= append(numbers, n)
		}

		for i:=0; i<len(numbers); i++{
			number = numbers[i]
			for _,v := range number{
				if v == 48 {
					currentcounter++
				}
				if zerocounter <= currentcounter{
					zerocounter = currentcounter
					longest = number
				}
			} 
			currentcounter = 0
	}
	longestnum,_ = strconv.ParseInt(longest,10,0)
	fmt.Println(zerocounter,longestnum)
}
