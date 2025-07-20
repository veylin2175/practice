package main

import (
	"errors"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 6, 7, 8, 10, 11, 13, 14, 16, 17, 18, 19}

	var index int
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println(err)
		return
	}

	newList, err := delEl(list, index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newList)
}

func delEl(list []int, el int) ([]int, error) {
	if el < 0 || el > len(list) {
		return nil, errors.New("index out of range")
	}

	return append(list[:el], list[el+1:]...), nil
}
