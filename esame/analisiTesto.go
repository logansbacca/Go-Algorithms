package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var biggest int
	var median float64
	smallest := 109239432049324
	Vocabolario:= make(map[string]int)
	myFile, err := os.Open(os.Args[1])
	var r rune   
	// so pode converter char em rune entao tem que pegar o valor zero no iddice 2 de os.args
	r= rune(os.Args[2][0])
	if err != nil {
		fmt.Print(err)
	}
	defer myFile.Close()

	scanner := bufio.NewScanner(myFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if _, exists:= Vocabolario[word]; exists {
			Vocabolario[word] = Vocabolario[word] + 1
		} else {
			Vocabolario[word] = 1
		}
	}

	var keys [] string
	for k := range Vocabolario {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, ":", Vocabolario[k])
		if smallest > Vocabolario[k] {
			smallest = Vocabolario[k]
		}
		if biggest < Vocabolario[k] {
			biggest = Vocabolario[k]
		}
		median = median + float64(Vocabolario[k])
	}
	median = median / float64(len(Vocabolario))

	var rs [] string
	for _, k := range keys {
		if rune(k[0]) == r{
			rs = append(rs, k)
		}
	}
	fmt.Println(smallest,biggest,median)
	fmt.Println(rs)
	
}
