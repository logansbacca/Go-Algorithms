package main

import "fmt"

func main() {
	words := []string{"dino", "dog", "elephant", "rhyno", "penguin", "alpaca"}
	for i := 0; i < len(words)-1; i++ {
		if len(words[i]) > len(words[i+1]) {
			words[i], words[i+1] = words[i+1], words[i]
			i = -1
		}
	}
	fmt.Print(words)
}
