package main

import "fmt"
import "strings"
import "os"

func main() {
	input := os.Args[1:]
	fmt.Println(maxConsonanti(input))
}

func maxConsonanti(lista []string) int {
	var max int
	var count int
	var p string
	for _, parola := range lista {
		for _, c := range parola {
			if strings.IndexRune(" bcdfghjklmnpqrstvwxyz ", c) != -1 {
				count++
			}
		}
		if max<count{
			max= count
			p= parola
		}
		count = 0
	}
	fmt.Print("la paraola con più consonanti è " ,p, " " )
	return max
}

