package main

import "fmt"

func main() {
	var numero int
	count:=1
	fmt.Scan(&numero)
	max := numero
	for i:= 1; i < 3; i++ {
		fmt.Scan(&numero)
		if numero > max {
			max = numero
		} else if numero == max{
			count++
		}
	}
	fmt.Println(max)
	fmt.Println(count)
}
