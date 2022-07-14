/* - Scrivere una funzione isPalindrome(s string) bool. La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma (è la stessa stringa sia letta da sinistra a destra che letta da destra a sinistra), e false altrimenti. */

package main

import "fmt"

func main() {
	fmt.Print(isPalindrome("logan"))
}

func isPalindrome(s string) bool {
	rev := ""
	for i:=len(s)-1; i>=0; i-- {
		rev = rev + string(s[i])
	}
	if rev != s {
		return false
	}
	return true
}