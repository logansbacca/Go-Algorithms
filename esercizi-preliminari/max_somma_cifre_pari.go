
package main
import "fmt"

func main() {
    var input int
    sommaPariMax := 0
    numeroSommaMax := 0
    
    for {
        
        fmt.Scanf("%d",&input)
        
        if input == -1 {
            break
        }
        
        sommaPari := 0
        preInput := input
        for input != 0 {
            
            cifra := input % 10
            
            if cifra % 2 == 0 {
                sommaPari += cifra
            }
            
            input /= 10
            
        }
        
        // sommaPari somma cifre pari di input
        if sommaPari > sommaPariMax {
            sommaPariMax = sommaPari
            numeroSommaMax = preInput
        }
        
    }
	fmt.Print(numeroSommaMax)
}
