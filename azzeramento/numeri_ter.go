package main

import "fmt"
func main() {
	var n int
	x := 0
	tot := 0
	for n = 0; n < 10; n++ {
		fmt.Scan(&n)
		if n == 0 {
			continue
		}
		x += n 
		tot++
	}
	fmt.Println(x)
	fmt.Println(tot)
}