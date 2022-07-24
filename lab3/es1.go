/* es12.go: genera una serie di valori casuali tra 0 e 20, stampandoli, fino a generare il valore sentinella K=20, e stampa quanti ne ha generati (20 escluso).

[aggiunto esempio di sentinella sbagliata con float saltato]

----
Nota.
Usare la funzione rand.Int() (o rand.Intn(n int)) del package "math/rand".
Per generare sequenze sempre diverse, importare il package "time" e usare l'istruzione rand.Seed(time.Now().Unix()) prima di iniziare a generare numeri.
E se i numeri da generare fossero tra 1 e 10, come faccio? */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generates new time each time
	rand.Seed(time.Now().Unix())
	var rn int
	var counter int
	
	for {
		if rn == 20 {
			break
		}
		rn= rand.Intn((21-0)+0)
		counter++
	}
	fmt.Print(counter)
}
