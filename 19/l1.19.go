package main

import (
	"fmt"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}
	runes := []rune(str)

	i := 0
	j := len(runes) - 1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	reverseStr := string(runes)
	fmt.Println(reverseStr)
}
