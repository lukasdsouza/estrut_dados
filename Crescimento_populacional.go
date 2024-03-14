package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scan(&PA, &PB, &G1, &G2)
		tempo := 0.0
		for PA <= PB {
			PA += float64(PA) * G1 / 100
			PB += float64(PB) * G2 / 100
			tempo++
			if tempo > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}
		if tempo <= 100 {
			fmt.Printf("%.0f anos.\n", tempo)
		}
	}
}