package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	longest := ""
	for i := 0; i < len(input); i++ {
		if len(longest) <= len(input[i]) {
			longest = input[i]
		}
	}
	fmt.Println(longest)
}
