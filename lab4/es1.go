/* - legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: 5Nov.),
- per ciascun carattere letto stabilisce se è una lettera minuscola (es. f)
 e stampa un messaggio per riga:
 - se minuscola, il carattere e la sua posizione nell'alfabeto (es "f è la 6^a");
 - altrimenti "altro"
- quando termina, stampa "bye" */

package main

import "fmt"

func main() {
	var char byte
	var alf = "abcdefghijklmnopqrstuvz"
	var index int

	for {
		if char == '.' {
			break	
		}
		fmt.Scanf("%c", &char)
		if char >= 'a' && char <= 'z' {
			for i, v := range alf {
				if byte(v) == char {
					index = i + 1
				}
			}
			fmt.Println(string(char), "è la", index,"^ a")
		} else if char == '.'{
			fmt.Println("bye")
		} else {
			fmt.Println("altro")
		}
	}
}
