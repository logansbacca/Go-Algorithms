/* Scrivere un programma disegna_slash.go che legge un intero
positivo n e stampa un back-slash (\) di asterischi di altezza n.
Esempio di esecuzione
dimensione \: 3
*
 *
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}
