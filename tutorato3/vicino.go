package main

import (
	"math"
	"fmt"
)

func main () {
	a := [] int {17,12,8,19}
	b:= 4
	fmt.Print(numeroVicino(a, b))
}


func numeroVicino(numeri []int, target int) int {
	closest := numeri[0]
	for i := 0; i < len(numeri); i++ {
		if math.Abs(float64(target-numeri[i])) < math.Abs(float64(target-closest)) {
			closest = numeri[i]
		}
	}
	return closest
}
