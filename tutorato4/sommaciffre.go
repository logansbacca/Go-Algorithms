/* Specifiche: Scrivere un programma che legge 10 interi positivi e stampa l’intero la cui somma delle cifre è
maggiore. Ad esempio, se la sequenza letta è 1 5 14 27 111 234 87 14 6 1221 il programma stampa
87 (la cui somma delle cifre è 15).
Progettazione: Completate ciascuna affermazione scegliendo l’opzione corretta. */

package main

import "fmt"

func main() {
	var num int
	var lastdigit int
	var sum int
	var start int
	var neg bool
	res := []int{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		//checking if num is negative
		if num < 0 {
			num = num * -1
			neg = true
		}
		// only working with double digit numbers
		if num > 9 {
			for num > 0 {
				lastdigit = num % 10
				res = append(res, lastdigit)
				num = num / 10
			}
			for i := 0; i < len(res); i++ {
				sum += res[i]
			}

			if sum > start {
				start = sum
				sum = 0

			}
			res = nil
		}

	}
	if neg {
		fmt.Print(start * -1)
	} else {
		fmt.Print(start)
	}

}
