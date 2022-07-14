package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	//"math"
)

func main() {
	length := len(os.Args[1:])
	sequenceParser(length)
	fmt.Println("")
}

func sequenceParser(length int) string {
	//a := ""

	epsilon, _ := strconv.ParseFloat(os.Args[length], 64)
	// convertendo string para float64

	for i := 1; i < length-1; i++ {
		current, _ := strconv.ParseFloat(os.Args[i], 64)
		next, _ := strconv.ParseFloat(os.Args[i+1], 64)

		dif := next - current
		dif = math.Abs(dif) // pego o valor absoluto/ sem sinal

		if dif > epsilon {
			if current > next {
				fmt.Println("+")
			}

			if current < next {
				fmt.Println("-")
			}

		} else {
			fmt.Println("=")
		}
	}

	return "a"
}

// 5.5 7.2 6.1 2.3 8.7 9.2 10.0 5 | 2
// 00+-00+
