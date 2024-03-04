package main

import "fmt"

func procuraMatriz(matrix [][]int, n int, x int) bool {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == x {
                return true
           	}
        }
    }
    return false
}

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    n := 3
    x := 0

    found := procuraMatriz(matrix, n, x)
    fmt.Printf("Value %d found: %t\n", x, found)
}