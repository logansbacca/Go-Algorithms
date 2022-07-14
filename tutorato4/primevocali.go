package main

import "fmt"
import "os"

func main () {
	lista := os.Args[1:]
	primeVocali(lista)
}

func primeVocali ( lista []string ) {
	for _ , parola := range lista {
	res := false
	for _ , v := range parola {
		if v == 'a' || v == 'e' || v=='i'|| v=='o' || v=='u'{
			res = true
			fmt.Println(string(v))
			break
		}
	}
	if !res  {
		fmt.Print("la parola non contiene vocali")
	}
	}
	}