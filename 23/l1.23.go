package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 6, 7, 8, 10, 11, 13, 14, 16, 17, 18, 19}

	var index int
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println(err)
		return
	}

	newList := delEl(list, index)
	fmt.Println(newList)
}

func delEl(list []int, el int) []int {
	newList := make([]int, len(list)-1)
	for k := range list {
		if k < el {
			newList[k] = list[k]
		} else if k != len(list)-1 && k >= el {
			newList[k] = list[k+1]
		} else {
			continue
		}
	}

	return newList
}
