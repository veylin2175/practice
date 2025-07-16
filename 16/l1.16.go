package main

import "fmt"

func main() {
	list := []int{5, 6, 2, 10, 1, 3, 9}

	fmt.Println("До сортировки:", list)
	list = quickSort(list)
	fmt.Println("После сортировки:", list)
}

func quickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[len(list)/2]
	var a, b []int
	for _, v := range list {
		if v < pivot {
			a = append(a, v)
		} else if v > pivot {
			b = append(b, v)
		}
	}

	a = quickSort(a)
	b = quickSort(b)

	res := append(a, pivot)
	res = append(res, b...)
	return res
}
