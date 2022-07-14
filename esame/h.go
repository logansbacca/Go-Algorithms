package main

import (
	"fmt"
)

func main() {
	fmt.Println(test(1, 2))
}

func test(a, b int) bool {
	if a == b {
		return true
	} else if a > b {
		return true
	}
	return false
}
