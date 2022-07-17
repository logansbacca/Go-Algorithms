/* scrivre un programma che legga da standard input un numero terminato da eof linea per liean se il numero dopo e piu grande mettere +, se piu piccolo - e se uguale mettere =  */

package main

import (
	"fmt"
	"io"
)

func main() {
	var err error
	var current int
	var previous int
	var res []string 
	fmt.Scan(&current)
	previous = current
	for {
		_, err = fmt.Scan(&current)
		if err == io.EOF {
			break
		}
		if current > previous {
			res = append(res, "+")
		}
		if previous > current {
			res = append(res, "-")
		}
		if previous == current {
			res = append(res, "=")
		}
		previous = current

	}
	res = append(res, "\n")
	for _, v := range res {
		fmt.Print(v)
	}

}
