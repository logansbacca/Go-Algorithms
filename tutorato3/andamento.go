package main

import "fmt"

func main() {
	var previous, current int

	fmt.Scan(&previous)

	for previous != 0 {
		fmt.Scan(&current)
		if current >= previous {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
		previous = current
	}
	return
}
