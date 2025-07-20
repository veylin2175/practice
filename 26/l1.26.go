package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qwertyuiop"

	unique := checkUnique(str)
	fmt.Println(unique)
}

func checkUnique(str string) bool {
	mapa := make(map[int32]bool)

	str = strings.ToLower(str)

	for _, v := range str {
		if mapa[v] {
			return false
		}
		mapa[v] = true
	}

	return true
}
