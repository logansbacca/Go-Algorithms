/* Scrivere un programma disegna_v.go che legge un intero
positivo n e stampa una v di asterischi di altezza n.
Esempio di esecuzione
dimensione v: 3
        
*   *   3
 * *    1
  *

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 3
	b := n
	for x := 0; x < n-1; x++ {
		fmt.Print(strings.Repeat(" ", x))
		fmt.Print("*")
		fmt.Print(strings.Repeat(" ", b))
		fmt.Print("*")
		fmt.Println()
		if b > 1 {
			b = b - 2
		}
	}

	fmt.Print(strings.Repeat(" ", n-1))
	fmt.Print("*")

}