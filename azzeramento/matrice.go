package main

import "fmt"

func main() {
	const R, C = 3, 4
	var a [R][C]int
	fmt.Print(a)
	for i := 0; i < R; i++ { 
		for j := 0; j < C; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			a[i][j] *= 2
		}
	}
	fmt.Println(a)
}

/* 1. Cosa indica la variabile i?
	i indica la riga corrente
2. A cosa serve il ciclo alle righe 15-19?
	moltiplica ogni numero su ogni riga per 2
3. Il ciclo for alle righe 15-17 può essere sostituito con un ciclo for range?
Si puo essere sostitutito con un ciclo for range.
4. Riassumete in una frase cosa fa il programma.
il programma itera su tutte le righe e colonne della matrice e assegna in input il valore dello scan. Dopo di che in un altro loop, moltiplica ogni numero in ogni collona e righa della matrice per 2.
\*/
