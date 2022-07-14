// given a word, change each vowel present to *

package main

import "strings"
import "fmt"

func main () {
	str := [] string {}
	var x string = "nicolò è una persona bella"

	for _,v := range x {
		a:=  strings.ToLower(string(v))
		if (a == "ò" || a == "è" || a == "é"|| a == "ù"|| a == "ì") || a == "a"|| a == "e"|| a == "i"|| a == "o"|| a == "u"{
			str = append(str, "*")
		} else {
			str = append(str, string(v))
			
		}
		
	}
	for _,v := range str {
		fmt.Print(v)
	}
	
}