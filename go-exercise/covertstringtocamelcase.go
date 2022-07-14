package main

import (
	"fmt"
	"strings"
)

func main() {
	ToCamelCase("the_stealth-warrior") // this should return theStealthWarrior
}

func ToCamelCase(s string) string {

	a := strings.Replace(s, "-", "", -1)
	b := strings.Replace(s, "_", "", -1)
	fmt.Print(a, b)
	return a
}
