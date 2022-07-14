/* Scrivere un programma disegna_v.go che legge un intero
positivo n e stampa una v di asterischi di altezza n.
Esempio di esecuzione
dimensione v: 3
* *
 * *
 * */

package main
import "fmt"

func main ()  {
	var num int
	fmt.Scan(&num)
	for i := 0; i < num-1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for k :=0 ; k < (num - i - 2) + (num - i - 1); k++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
	for i:=0; i<num-1; i++ {
		fmt.Print(" ")
	}
	fmt.Println("*")
}


