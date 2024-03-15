package main

import "fmt"

func anosParaUltrapassar(PA, PB int, G1, G2 float64) string {
	anos := 0
	for PA <= PB {
		PA = int(float64(PA) * (1 + G1/100))
		PB = int(float64(PB) * (1 + G2/100))
		anos++
		if anos > 100 {
			return "Mais de 1 seculo."
		}
	}
	return fmt.Sprintf("%d anos.", anos)
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scan(&PA, &PB, &G1, &G2)
		fmt.Println(anosParaUltrapassar(PA, PB, G1, G2))
	}
}
