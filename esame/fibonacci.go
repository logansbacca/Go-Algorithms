/* Scrivere un programma fibonacci.go che legge un intero
positivo n e stampa i numeri di fibonacci dal primo all'n-esimo,
rappresentandoli come righe di asterischi, ciascuna lunga quanto
il numero da rappresentare.
Esempio di esecuzione
un numero: 6
*
*
**
***
*****
******** */

package main

import "fmt"

func main() {
	fib1 := 0
	fib2 := 1
	for i:=1; i <= 6; i++ {
		fib1, fib2 = fib2, fib1+fib2
		for j:= 1; j<=fib1;j++ {
			fmt.Print("*")
		}
		fmt.Println()
		}
		
	}



