package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	pariDispari(numeri)
}


func pariDispari(numeri []int)  {
	for i:=0; i<len(numeri);i++ {
		numero := numeri[i]
		if numero % 2 == 0  {
			fmt.Println("pari")
		}else {
			fmt.Println("dispari")
		}
	}
	return 
}