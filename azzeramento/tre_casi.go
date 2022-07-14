package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	if len(os.Args) != 4 {
		fmt.Println(" lanciare il programma con tre argomenti ")
		return
	}
	fmt.Println(" Inserisci un numero :")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(" ci vuole un numero intero ")
		return
	}
	switch n {
	case 1:
		os.Args[2], os.Args[3] = os.Args[3], os.Args[2]
	case 2:
		os.Args[1], os.Args[3] = os.Args[3], os.Args[1]
	case 3:
		os.Args[1], os.Args[2] = os.Args[2], os.Args[1]
	}
	fmt.Println(os.Args[1], os.Args[2], os.Args[3])
}
