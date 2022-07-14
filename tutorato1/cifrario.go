package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var msg string
	fmt.Println("Enter Message :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg = scanner.Text()
	encode(msg)
}

func encode(str string) {
	len := len(str)
	newmsg := ""
	str = strings.ToUpper(str)
	var key int
	fmt.Println("Insert key :")
	fmt.Scanln(&key)
	for i := 0; i < len; i++ {
		lettera := str[i]
		if lettera >= 'A' && lettera <= 'Z' {
			posizione := lettera - 'A'
			nuovaposizione := (rune(posizione) + rune(key)) % 26
			newmsg += string('A' + nuovaposizione)
		} else {
			newmsg += string(lettera)
		}
	}
	fmt.Print(strings.ToLower(newmsg))
}
