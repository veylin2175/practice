package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a    int
		str  string
		flag bool
		ch   chan int
	)
	getType(a)
	getType(str)
	getType(flag)
	getType(ch)

	// 2 способ
	getType2(a)
	getType2(str)
	getType2(flag)
	getType2(ch)
}

func getType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t)
}

// Альтернатива через switch v.(type)

func getType2(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	default:
		fmt.Println("Неиспользуемый в свитче тип")
	}
}
