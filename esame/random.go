package main

import "fmt"
import "math/rand"
import "time"


func main () {
	rand.Seed(time.Now().UnixNano())
	random:= 0 + rand.Intn(100-0)
	fmt.Print(random)
}