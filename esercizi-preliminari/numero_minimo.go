package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := string(scanner.Text())
	outputNumber := []rune{}

	for i := 0; i < len(input); i++ {
		input = strings.ReplaceAll(input, " ", "")
		currentDigit := rune(input[i])
		outputNumber = append(outputNumber, currentDigit)

	}

	for x := 0; x < len(outputNumber)-1; x++ {
		if outputNumber[x] > outputNumber[x+1] {
			outputNumber[x], outputNumber[x+1] = outputNumber[x+1], outputNumber[x]
			x = -1
		}
	}
	fmt.Print(string(outputNumber[0]))
}
