package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(minDispari(numeri))
}


func minDispari(numeri []int) int  {
	minnum:= numeri[0]
	for i:=0; i<len(numeri);i++ {
		if numeri[i] % 2 !=0  && numeri[i] < minnum {
			minnum = numeri[i]
		}
	}
	return minnum
}