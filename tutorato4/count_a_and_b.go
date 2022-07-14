package main
import "fmt"
func main () {
var parola string
var counta int
var countb int

for i:=0; i<10; i ++ {
fmt.Scan(&parola )
if parola[0] == 'a' {
	counta ++
}
if parola[0] == 'b' {
	countb++
}
}

fmt .Println( counta )
fmt .Println( countb )
}
