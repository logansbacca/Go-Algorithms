package main
import "fmt"
import "os"
import "strconv"

func main() {
	stop := 9999
	const FATTORE_SCALA, ZERO = 5.0 / 9.0, 32.0
	if len(os.Args) < 2 {
		fmt.Println(" Nessun dato in input ")
		return
	}
	for _, temp := range os.Args[1:] {
		intval, _ := strconv.Atoi(temp)
		if intval == stop {
			break
		}
		fahr, _ := strconv.ParseFloat(temp, 64)
		fmt.Println((fahr - ZERO) * FATTORE_SCALA)
	}
}
