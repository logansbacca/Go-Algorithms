/* Scrivere un programma pos_alfabeto.go che
- legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: 5Nov.),
- per ciascun carattere letto stabilisce se è una lettera minuscola (es. f)
 e stampa un messaggio per riga:
 - se minuscola, il carattere e la sua posizione nell'alfabeto (es "f è la 6^a");
 - altrimenti "altro"
- quando termina, stampa "bye"
 */

 package main

 import "fmt"
 import "unicode"
 import "strings"

 func main () {
	var c byte
	for  {
		if c == '.' {
			break
		}
		fmt.Scanf("%c", &c)
		if unicode.IsLower(rune(c)) {
			fmt.Print(string(c), " è la ", strings.IndexRune("abcdefghijklmnopqrstuvwxyz", rune(c)) + 1, "^a\n")
		} else if unicode.IsUpper(rune(c)){
			fmt.Print("alto\n")
		}
	}
	fmt.Print("bye")
 }