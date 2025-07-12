package main

import (
	"fmt"
	"strconv"
)

func SetBit(num int64, i uint, bit uint) int64 {
	if bit == 1 {
		return num | (1 << i)
	}
	return num &^ (1 << i)
}

func main() {
	var num int64
	var i uint
	var bit uint

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ожидается число...", err)
		return
	}

	fmt.Print("Введите номер бита (0-63): ")
	_, err = fmt.Scan(&i)
	if err != nil || i > 63 {
		fmt.Println("Ошибка: номер бита должен быть от 0 до 63")
		return
	}

	fmt.Print("Введите значение бита (0 или 1): ")
	_, err = fmt.Scan(&bit)
	if err != nil || (bit != 0 && bit != 1) {
		fmt.Println("Ошибка: значение бита должно быть 0 или 1")
		return
	}

	result := SetBit(num, i, bit)

	fmt.Printf("\nИсходное число: %d (%064s)\n", num, strconv.FormatInt(num, 2))
	fmt.Printf("Установка %d-го бита в %d\n", i, bit)
	fmt.Printf("Результат:      %d (%064s)\n", result, strconv.FormatInt(result, 2))
}
