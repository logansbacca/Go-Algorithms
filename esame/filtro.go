package main

import "fmt"

func main() {
	var parola string
	var carattere string
	var counter int
	fmt.Scan(&parola)
	fmt.Scan(&carattere)

	for i:=0; i<len(parola); i++{
		if string(parola[i]) == carattere {
			counter++
		}
	}
	fmt.Print(counter)


}
