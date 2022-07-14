package main

import "fmt"
func main() {
	var n int
	x := 0
	tot := 0
  media := 0
	for n = 0; n < 10; n++ {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		x += n 
		tot++
	}
	media = x/tot
  fmt.Print(media)
}
