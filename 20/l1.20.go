package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(trueRealise(str))

	str = "snow dog sun"
	fmt.Println(easyRealise(str))
}

// trueRealise меняет слова местами согласно условию (без использования доп слайса)
// список rune - исключение, так как строки в Go неизменяемы
func trueRealise(str string) string {
	runes := []rune(str)

	i := 0
	j := len(runes) - 1

	// сначала меняем порядок всех символов
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]

		i++
		j--
	}

	// далее - меняем символы в каждом слове по отдельности обратно
	n := 0
	m := len(runes)
	for i := 0; i < m; i++ {
		if runes[i] == ' ' || i == m-1 {
			var temp int
			if i == m-1 {
				temp = i
			} else {
				temp = i - 1
			}

			for n < temp {
				runes[n], runes[temp] = runes[temp], runes[n]
				n++
				temp--
			}
			n = i + 1
		}
	}

	return string(runes)
}

// easyRealise - простая функция, просто разделяет слова
// и меняет их местами, но с использованием доп слайса
func easyRealise(str string) string {
	list := strings.Split(str, " ")

	i := 0
	j := len(list) - 1

	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}

	str = strings.Join(list, " ")

	return str
}
