package main

import "fmt"

func acharPar(vetor []int, alvo int) (int, int) {
	mapa := make(map[int]bool)
	for _, numero := range vetor {
		complemento := alvo - numero
		if mapa[complemento] {
			return numero, complemento
		}
		mapa[numero] = true
	}
	return -1, -1
}

func main() {
	vetor := []int{1, 5, 3, 8, 4}
	alvo := 7

	par := acharPar1(vetor, alvo)
	fmt.Printf("Par encontrado: (%d, %d)\n", par[0], par[1])
}