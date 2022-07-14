package main

import "fmt"

func main() {
	var previous, current int
	count := 0
	fmt.Scan(&previous)
	count++
	for {
		fmt.Scan(&current)
		if current < previous {
			break
		}
		previous = current
		count++
	}
	fmt.Println(" Numeri validi inseriti :", count)
}
