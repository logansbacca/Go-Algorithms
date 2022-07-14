package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	fmt.Println(piucorta(parole))
}


func piucorta (parole [] string) int{
	piccola := parole[0]
	for i:=0;i<len(parole); i++ {
		parolacorrente:= parole[i]
		if parolacorrente < piccola {
			piccola = parolacorrente
		}
	}
	return len(piccola)
}