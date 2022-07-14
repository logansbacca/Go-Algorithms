package main

import "fmt"

//https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

func main() {
	ages := []int{10, -22, 3, 55, -123, 71, 15, 100, -8, 3}
	bigNum := 0
	for i := 0; i < len(ages); i++ {
		if ages[i] < ages[i+1] {
			bigNum = ages[i]
			break

		}

	}
	fmt.Println(bigNum)
}
