package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 7, 8}
	b := []int{2, 3, 4, 5, 6}

	crossed := crossArr(a, b)
	fmt.Println(crossed)
}

func crossArr(a []int, b []int) []int {
	exists := make(map[int]bool, len(a))
	for _, v := range a {
		exists[v] = true
	}

	crossed := make([]int, 0, 10)
	for _, v := range b {
		if exists[v] {
			crossed = append(crossed, v)
		}
		delete(exists, v)
	}

	return crossed
}
