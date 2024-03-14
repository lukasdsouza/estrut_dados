package main

import "fmt"

func mergeSortedLists(a []int, m int, b []int, n int) []int {
 v := make([]int, 0, m+n)
 i, j := 0, 0
 for i < m || j < n {
 	if i < m && j < n {
 		if a[i] < b[j] {
 			v = append(v, a[i])
 			i++
 		} else {
 			v = append(v, b[j])
 			j++
 		}
 	} else if i < m {
 		v = append(v, a[i:m]...)
 		i = m
 	} else {
 		v = append(v, b[j:n]...)
 		j = n
 	}
 }
 return v
}

func insertionSort(arr *[]int) {
 for i := 1; i < len(*arr); i++ {
 	key := (*arr)[i]
 	j := i - 1
 	for j >= 0 && (*arr)[j] > key {
 		(*arr)[j+1] = (*arr)[j]
 		j--
 	}
 	(*arr)[j+1] = key
 }
}

func main() {
 a := []int{1, 3, 5}
 b := []int{2, 4, 6}
 v := mergeSortedLists(a, len(a), b, len(b))
 insertionSort(&v)
 fmt.Println(v)
}