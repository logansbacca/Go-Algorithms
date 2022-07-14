package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}

	fmt.Println(piuCorta(parole))

}

func piuCorta(parole []string) int {
	smallest := 0

	for i := 0; i < len(parole)-1; i++ {
		if parole[i] < parole[i+1] {
			parole[i], parole[i+1] = parole[i+1], parole[i]
			i = -1
		}
	}

	a := len(parole[0])
	totsize := (a * 8) / 8
	smallest = totsize
	return smallest

}
