package main

import "fmt"

// Função recursiva para resolver o problema da Torre de Hanói
func hanoi(n int, source, auxiliary, target int) {
	if n > 0 {
		hanoi(n-1, source, target, auxiliary)
		fmt.Printf("Mover disco %d de haste %d para haste %d\n", n, source, target)
		hanoi(n-1, auxiliary, source, target)
	}
}

// Função para encontrar um par de números no array que somados resultem em um alvo dado
func findPair(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == target {
			return arr[left], arr[right]
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

func main() {
	const nDisks = 3
	hanoi(nDisks, 1, 2, 3)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 12
	pair, pair2 := findPair(arr, target)
	fmt.Println(pair, pair2)
}
