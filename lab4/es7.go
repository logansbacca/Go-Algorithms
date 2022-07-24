/* Scrivere un programma disegna_slash.go che legge un intero
positivo n e stampa un back-slash (\) di asterischi di altezza n.
Esempio di esecuzione
dimensione \: 3
*
 *
  *
*/

package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scan(&n)
	/* for i := 0; i < n; i++ {
		for j:= 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	} */

	for x:=0; x<n; n++{
		strings.Repeat(" ",x)
		fmt.Println()
	}
	fmt.Print("*")
	


}
