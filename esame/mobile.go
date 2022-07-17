package main

import "fmt"

type Utenza struct {
	NUMERO_TELEFONO int
	CODICE_SIM      rune
}

func main() {
	LeggiUtenze()
}

func LeggiUtenze() (utenze []Utenza) {

	for {	
		_,err := fmt.Scan(&utenze)
		if err != nil {
			break
		}
	}
	fmt.Print(utenze)
	return utenze
}
