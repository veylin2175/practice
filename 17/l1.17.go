package main

import "fmt"

func main() {
	list := []int{2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14, 15, 17, 18, 21, 23, 24, 27, 28, 29, 31, 34, 35, 37, 38, 39}
	var el int
	_, err := fmt.Scan(&el)
	if err != nil {
		fmt.Println(err)
		return
	}

	index := biSearch(list, el)
	if index == -1 {
		fmt.Printf("Элемент %v не найден", el)
	} else {
		fmt.Printf("Элемент %v находится по индексу %v", el, index)
	}
}

func biSearch(list []int, el int) int {
	left, right := 0, len(list)-1

	for left <= right {
		mid := left + (right-left)/2

		if list[mid] == el {
			return mid
		} else if list[mid] > el {
			right = mid - 1
		} else if list[mid] < el {
			left = mid + 1
		}
	}

	return -1
}
