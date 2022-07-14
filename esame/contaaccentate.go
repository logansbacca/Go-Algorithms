/* Scrivere un programma (contaAccentate.go) che, dDato un testo in input,
conta quante lettere accentate ci sono (considerate solo à, è, é, ì, ò, ù).
Leggere una runa alla volta, utilizzando il metodo Split per Scanner
del package bufio con la funzione di split bufio.ScanRunes:
myScanner := bufio.NewScanner(os.Stdin)
myScanner.Split(bufio.ScanRunes)

Esempio di esecuzione:

$ go run contaAccentate.go < file_uno
Accentate trovate:  6 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {
	var counter int
	Scanner:= bufio.NewScanner(os.Stdin)
	Scanner.Split(bufio.ScanRunes)
	
	for Scanner.Scan() {
		err := Scanner.Err()
		if err == io.EOF {
			break
		}
		rn := Scanner.Text()
		if rn == "à" || rn == "è" || rn == "ù" || rn == "ò" || rn == "ì" || rn == "é" {
			counter++
		}
}
	fmt.Println(counter)
}