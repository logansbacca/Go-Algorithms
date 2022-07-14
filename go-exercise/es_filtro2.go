package main

import (
	"fmt"
	"os"
)

func main() {
	inputNumber := []rune(os.Args[1])
	currentDigit := rune(65)
	outputNumber := []rune{}

	for i := 0; i < len(inputNumber); i++ {
		if inputNumber[i] != currentDigit {
			currentDigit = inputNumber[i]
			outputNumber = append(outputNumber, currentDigit)
		}
	}
	fmt.Println(string(outputNumber))
}
