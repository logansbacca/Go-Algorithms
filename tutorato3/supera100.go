package main 

import "fmt"

func main () {
var interi int
var contatore int

for interi > -1 {
	fmt.Scan(&interi)
	if interi > 100 {
		fmt.Print(interi)
		return
	}
	contatore = interi
}

if contatore <100 {
	fmt.Println("nessun numero maggiore")
}

}