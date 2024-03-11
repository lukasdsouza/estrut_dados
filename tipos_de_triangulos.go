package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Lê os valores de ponto flutuante A, B e C
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	// Coloque os valores em ordem decrescente
	lados := []float64{A, B, C}
	sort.Sort(sort.Reverse(sort.Float64Slice(lados)))

	A = lados[0]
	B = lados[1]
	C = lados[2]

	// Verifique o tipo de triângulo
	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		// Verifique se é um triângulo retângulo
		if math.Pow(A, 2) == math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO RETANGULO")
		}
		// Verifique se é um triângulo obtusângulo
		if math.Pow(A, 2) > math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		// Verifique se é um triângulo acutângulo
		if math.Pow(A, 2) < math.Pow(B, 2)+math.Pow(C, 2) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		// Verifique se é um triângulo equilátero
		if A == B && B == C {
			fmt.Println("TRIANGULO EQUILATERO")
		}
		// Verifique se é um triângulo isósceles
		if (A == B && B != C) || (A != B && B == C) {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
