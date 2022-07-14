package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi((os.Args[1]))
	b, _ := strconv.Atoi((os.Args[2]))
	n, _ := strconv.Atoi((os.Args[3]))

	sum := 0
	for i := a + 1; i <= b-1; i++ {
		if i%n == 0 {
			sum = sum + 1
		}
	}
	fmt.Println(sum)
}
