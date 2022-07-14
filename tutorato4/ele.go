package main

import "fmt"

func main() {
	var dueGiorniFa, ieri, oggi int
	x:= 1
	fmt.Scan(&ieri)
	fmt.Scan(&oggi)
	for i:= 0; i <= 1; i++ {
		dueGiorniFa = ieri
		ieri = oggi
		fmt.Scan(&oggi)
		x++
		fmt.Println(" oggi : ", oggi, " ieri : ", ieri, " dueGiorniFa : ", dueGiorniFa)
		if ieri < dueGiorniFa && ieri < oggi {
			fmt.Println(" il consumo nel giorno ", x, "e ‘ stato inferiore ",
				" rispetto sia al giorno prima che al giorno dopo ")
				break
		}
		
	}
}
