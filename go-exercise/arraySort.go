package main

import "fmt"

func main() {
	a := []int{99, 1, 16, 15}
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
		} else {
			a[i], a[i+1] = a[i+1], a[i]
			i = -1
		}
	}
	fmt.Print(a[len(a)-1], a[len(a)-2])
}
