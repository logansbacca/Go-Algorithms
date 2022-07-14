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
	res := 1
	for i := 0; i < len(numeri); i++ {
		if numeri[i] > 7 && numeri[i]%4 == 0 {
			res = res * numeri[i]
		}
	}
	fmt.Print(res)
	return res

}
