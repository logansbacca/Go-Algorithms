package main

import "fmt"

func main () {
	stop:= 0
	var num int
	var numpari [] int
	numpari = make([]int, 0)

	for {
		fmt.Scan(&num)
		if num == stop {
			break
		}
		if num %2 == 0 {
			numpari = append(numpari, num)
		}
	}
	tot:= 0
	for i:=0; i<len(numpari) -1;i++ {
		tot = numpari[i] * numpari [i+1]
	}
	fmt.Print(tot)
}