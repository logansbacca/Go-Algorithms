/* scrivere un programma (lineePariDispari.go) che, dato un testo in input,
stampa prima tutte le righe pari e poi tutte le dispari.
La prima riga del testo è la riga 1.

Esempio di esecuzione:

$ go run lineePariDispari.go < file_uno
cominciò a gridar la fiera bocca,
E ’l duca mio ver lui: «Anima sciocca,
quand’ira o altra passion ti tocca!
"Raphél maì amèche zabì almi",
cui non si convenia più dolci salmi.
tienti col corno, e con quel ti disfoga  */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line string
	var dispari []string
	var pari []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			break
		}
		line = scanner.Text()
		if len(line)%2 == 0 {
			pari = append(pari, line)
		}
		dispari = append(dispari, line)
	}
	for _, v := range pari {
		fmt.Println("pari")
		fmt.Println(v)
	}
	for _, v := range dispari {
		fmt.Println("dispari")
		fmt.Println(v)
	}
}
