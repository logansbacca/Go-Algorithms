package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	EPSILON, _ := strconv.ParseFloat(input[0],64)

	for i := 1; i < len(input)-1; i++ {
		x, _ := strconv.ParseFloat(input[i],64)
		y, _ := strconv.ParseFloat(input[i+1],64)
		res := math.Abs(x-y)
	
		if res > EPSILON {
			if x > y {
				fmt.Print("<")
			}

			if x < y {
				fmt.Print(">")
			}
			
		} else {
			fmt.Print("=")
		}
					
	}
}
