package main

import "os"
import "unicode"
import "fmt"

func main() {
	input := os.Args[1:]
	result := true
	for i:=0;i<len(input);i++ {
		letter:= input[i]
		for j:=0; j<len(letter)-1;j++{
			current := rune(letter[j])
			next := rune(letter[j+1])
			if unicode.IsLower(current) && !unicode.IsUpper(next) || unicode.IsUpper(current)  && !unicode.IsLower(next){
				result = false
				break
			} 
		}
	}
	if result == true {
		fmt.Print("sequenza valida")
	} else {
		fmt.Print("sequenza non valida")
	}
}
