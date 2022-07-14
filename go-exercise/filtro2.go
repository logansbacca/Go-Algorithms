package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := string(scanner.Text())
	currentDigit := "A" // incializar a variavel com algo diferente de numbero
	outputNumber := []string{}
	fmt.Print(len(input))

	for i := 0; i < len(input); i++ {
		if string(input[i]) == currentDigit {

		} else {
			currentDigit = string(input[i])
			outputNumber = append(outputNumber, string(input[i]))
		}
	}
	for i := 0; i < len(outputNumber); i++ {
		fmt.Print(outputNumber[i])
	}
}
