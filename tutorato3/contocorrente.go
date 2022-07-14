package main

import "fmt"

func main () {
var saldo int
var newspesa int
var totspese int

fmt.Println("inserire saldo")
fmt.Scan(&saldo)

fmt.Println("inserire spese")
for saldo>= 0 {
	fmt.Scan(&newspesa)
	totspese = newspesa
	saldo = saldo - totspese
}
fmt.Println(saldo)
}