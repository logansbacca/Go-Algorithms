package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}

	primacona(parole)
}


func primacona (parole [] string) {

	for i:=0;i<len(parole); i++ {
		parola := parole[i]
		for j:=0; j<len(parola); j++{
			if string(parola[j]) == "a"|| string(parola[j]) == "A"{
				fmt.Print(string(parola))
				return
			}
		}
	}
}