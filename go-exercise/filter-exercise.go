// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := 366
	convert := strconv.Itoa(input)
	for i := 0; i < len(convert[1:]); i++ {
		if convert[i] == convert[i+1] {
			strings.Replace(string(convert[i]), " ", ",", -1)
			fmt.Println(convert)
		} else {
			fmt.Println("no equal numbers")
		}
	}
}
