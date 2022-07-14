package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}

	fmt.Println(quantecona(parole))
}


func quantecona (parole [] string) int {
	var counter = 0
	for i:=0;i<len(parole); i++ {
		parola := parole[i]
		for j:=0; j<len(parola); j++{
			if string(parola[j]) == "a"|| string(parola[j]) == "A"{
				counter++
				break
			}
		}
	}
	return counter
}