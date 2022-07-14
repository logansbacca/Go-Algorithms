package main

import "fmt"

func main () {
	var char byte
	var word string
	fmt.Scanf("%c", &char)
	fmt.Printf("%c %c %c\n", char-1, char, char+1)

	if char >= 'A' && char <= 'L'{
		fmt.Println("A-L") 
	}  else {
		fmt.Println("altro")
	}

	fmt.Scan(&word)

	for _,c := range word {
		fmt.Println(string(c))
	}

		
		
}