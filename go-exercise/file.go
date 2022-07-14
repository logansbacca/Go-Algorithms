package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	myarray := []float64{3, 4, 5, 6, 2}
	v := sequenceParser(myarray)
	res := []string{"hello", "hi"}
	conv, _ := strconv.ParseFloat(res[1], 64)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(conv))
}

func sequenceParser(myarray []float64) string {
	lastnum := myarray[len(myarray)-1]
	result := ""
	for i := 0; i < len(myarray)-2; i++ {
		difference := myarray[i] - myarray[i+1]
		absolutevaldif := math.Abs(difference)
		if absolutevaldif > lastnum {
			if difference < 0 {
				result = result + "-"
			} else {
				result = result + "+"
			}
		} else {
			result = result + "0"
		}
	}
	return result

}
