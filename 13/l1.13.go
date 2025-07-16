package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("До: a: %v, b: %v\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После: a: %v, b: %v\n", a, b)
}
