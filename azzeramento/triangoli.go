package main 
import "fmt" 
func main () { 
var a , b , c int 
fmt.Println(" Inserisci tre numeri separati da spazi :") 
fmt.Scan (&a,&b,&c) 
if a == b && a == c { 
fmt .Println(" Il triangolo e isoscele") 
} else if a == b || a == c || b == c { 
fmt .Println(" Il triangolo e equilatero ") 
} else { 
fmt .Println(" Il triangolo e scaleno") 
} 
}
