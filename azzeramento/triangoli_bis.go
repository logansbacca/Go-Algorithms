package main 
import "fmt" 
func main () { 
var a , b , c int 
fmt.Println(" Inserisci tre numeri separati da spazi :") 
fmt.Scan (&a,&b,&c) 
//condizione di esistenza
if a>c+b && b> a+c && c> a+b{
	return
}
if a == b && a == c { 
fmt .Println(" Il triangolo e equilatero") 
} else if a == b || a == c || b == c { 
fmt .Println(" Il triangolo isoscele ") 
} else { 
fmt .Println(" Il triangolo e scaleno") 
} 
}
