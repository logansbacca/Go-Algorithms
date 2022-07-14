package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(stranoProdotto(numeri)) 
}


func stranoProdotto(numeri []int) int {
	result := 1
	for i:=0; i<len(numeri);i++ {
		numero := numeri[i]
		if numero % 4 == 0 && numero > 7 {
			result =  result * numero
		}
	}
	return result
}