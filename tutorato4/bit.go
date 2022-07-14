package main

import "fmt"

func main() {
	var bit, bitPrev, leng int
	fmt.Scan(&bit)
	const N = 8
	leng = 1
	for i := 1; i < N; i++ {
		bitPrev = bit
		fmt.Scan(&bit)
		if bit == bitPrev {
			leng++
		} else {	
			fmt.Println(leng)	
			leng = 1
		}
	}
	fmt.Print(leng)
}
