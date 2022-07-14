package main

import "fmt"

const STOP = 9999

func main() {
	var dayRain, totRain, count int
	fmt.Print(" mm di pioggia : ")
	var totdayRain []int
	totdayRain = make([]int, 0)
	for {
		fmt.Scan(&dayRain)
		if dayRain == STOP {
			break
		}
		totdayRain = append(totdayRain, dayRain)
		count++
		totRain += dayRain
	}

	mycount := 0

	if count == 0 {
		fmt.Println(" nessun dato disponibile ")
	} else {
		averageRainfall := float64(totRain) / float64(count)
		raincount := 0
		highprecipitation := 0
		for i := 0; i < len(totdayRain); i++ {
			if totdayRain[i] > int(averageRainfall) {
				highprecipitation = totdayRain[i]
			}
			if totdayRain[i] > raincount {
				raincount = totdayRain[i]
			}
		}
		fmt.Println("il giorno con la precipitazione piu alta della media è :", highprecipitation)
		fmt.Println("l'ultimo giorno di pioggia è :", raincount)
		fmt.Println(" media :", averageRainfall)
		if dayRain > int(averageRainfall) {
			mycount++
		}
	}
}
