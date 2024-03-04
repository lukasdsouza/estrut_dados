package main

import (
	"fmt"
)

func main() {
	var X1, Y1, X2, Y2, N int
	for {
		fmt.Scanf("%d %d %d %d", &X1, &Y1, &X2, &Y2)
		if X1 == 0 && Y1 == 0 && X2 == 0 && Y2 == 0 {
			break
		}
		fmt.Scanf("%d", &N)
		meteorites := make([][2]int, N)
		for i := 0; i < N; i++ {
			fmt.Scanf("%d %d", &meteorites[i][0], &meteorites[i][1])
		}
		count := 0
		for _, meteorite := range meteorites {
			if isInside(X1, Y1, X2, Y2, meteorite) {
				count++
			}
		}
		fmt.Printf("Teste %d\n%d\n\n", X1, count)
	}
}

func isInside(X1, Y1, X2, Y2 int, meteorite [2]int) bool {
	return X1 <= meteorite[0] && meteorite[0] <= X2 && Y1 <= meteorite[1] && meteorite[1] <= Y2
}
