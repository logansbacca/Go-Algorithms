//Scrivere un programma che legga da standard input una sequenza di numeri terminata da
//zero e conti quante sono le coppie di numeri naturali consecutivi uguali.

package main

import "fmt"

func main() {
	var numeroattuale int
	var numeroprecedente int
	fmt.Scan(&numeroattuale)
	var coppie int

	for {
		numeroprecedente = numeroattuale
		fmt.Scan(&numeroattuale)
		if numeroattuale == numeroprecedente {
			coppie++
		}
		if numeroattuale == 0{
			break
		}
	}
	fmt.Println(coppie)
}
