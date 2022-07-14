// Scrivere un programma che legga una sequenza di interi terminata da zero e stampi il più
//piccolo intero letto e il numero di volte che compare all’interno della sequenza
package main
import "fmt"

func main () {

	var num, piccolo, counter int
	fmt.Scan(&piccolo)
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if piccolo > num  {
			piccolo = num
			counter++
		}	
	}
	fmt.Print(piccolo,counter)
}