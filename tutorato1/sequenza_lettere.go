package main

import "fmt"

func main() {
	var carattere, prec rune
	var cestino rune
	var contatore int
	var maxdif int
	fmt.Scanf("%c", &prec)
	contatore ++
	fmt.Scanf("%c", &cestino)
	for {
		fmt.Scanf("%c", &carattere)
		fmt.Scanf("%c", &cestino)
		if prec >= carattere {
			break
		}
		if int(carattere- prec) > maxdif {
			maxdif = int(carattere - prec)
		}
		prec = carattere
		if carattere == ' ' || carattere == '\n' {
			continue
		}
		contatore++
	}
	fmt.Println("massima differenza riscontrata:", maxdif )
	fmt.Println("lunghezza:", contatore )

}
