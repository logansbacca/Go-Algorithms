package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Utenza struct {
	NUMERO_TELEFONO int
	CODICE_SIM      string
}

func main() {
	LeggiUtenze()
}

func LeggiUtenze() (utenze []Utenza) {
	var line [] string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), ";")
		intnum,_ := strconv.Atoi(line[0])
		s := Utenza{NUMERO_TELEFONO: intnum, CODICE_SIM: line[1]}
		utenze = append(utenze, s)
		line= nil
	}

	for _,v := range utenze {
		fmt.Println(v)
	}
	return utenze
}
