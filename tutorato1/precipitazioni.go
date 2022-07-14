package main

import "fmt"

const STOP = 9999

func main() {
	var dayRain, totRain, count int
	fmt.Print(" mm di pioggia : ")
	for {
		fmt.Scan(&dayRain)
		if dayRain == STOP {
			break
		}
		count++
		totRain += dayRain
	}
	if count == 0 {
		fmt.Println(" nessun dato disponibile ")
	} else {
		averageRainfall := float64(totRain) / float64(count)
		fmt.Println(" media :", averageRainfall)
	}
}

/*Analizzate il codice sorgente e rispondete per iscritto alle seguenti domande.
1. Il programma precipitazioni.go attraverso quale canale riceve i dati da elaborare?
fmt.Scan
2. Di che tipo sono i dati ricevuti in input dal programma?
sono interi
3. È necessario memorizzare l’intera sequenza di dati ricevuti in input dal programma? Giustificate la risposta.
Non e necessaria
4. Cosa deve avvenire affinché il ciclo for termini?
deve essere dato in input il numero 9999
5. I dati ricevuti da input sono pronti da elaborare o è necessario manipolarli prima in qualche modo (calcolarne/derivarne/estrarne altri dati)?
Vanno trasformati in float64
6. Se si vuole un programma che stampi il numero di giorni in cui le precipitazioni sono state sopra la media e quale è stato l’ultimo giorno di pioggia, per quali delle domande qui sopra deve cambiare la risposta e come cambia?
bisogna memorizzare la sequenza ricevuta.
7. Modificate ora il programma in modo che stampi il numero di giorni in cui le precipitazioni sono state sopra la media e quale è stato l’ultimo giorno di pioggia, salvatelo con nome precipitazioni_bis.go e caricatelo sul sito di upload. \*/

