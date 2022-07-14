package main


import "fmt"

func main () {
	var currentinput int 
	var previousinput int
	var total int
	fmt.Println("insert first input")
	fmt.Scan(&previousinput)
	fmt.Println("insert the rest of the input")

	for  {
		fmt.Scan(&currentinput)
		total = currentinput - previousinput
		previousinput = currentinput
		if previousinput == 0 {
			break
		}
		fmt.Println("Differenza:" ,total)
	}
	
}