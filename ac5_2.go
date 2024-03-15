package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	lados := []float64{A, B, C}
	sort.Sort(sort.Reverse(sort.Float64Slice(lados)))

	A = lados[0]
	B = lados[1]
	C = lados[2]

	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {

		if math.Pow(A, 2) == math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO RETANGULO")
		}

		if math.Pow(A, 2) > math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}

		if math.Pow(A, 2) < math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}

		if A == B && B == C {
			fmt.Println("TRIANGULO EQUILATERO")
		}

		if (A == B && B != C) || (A != B && B == C) {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
