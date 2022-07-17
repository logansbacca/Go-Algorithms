package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	var str2 string
	var elem string
	var res string
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	res = str1
	for _, v := range str2 {
		elem = string(v)
		if strings.Contains(str1, elem) {

		} else {
			res = res + string(v)
		}
	}
	fmt.Print(res)
}
