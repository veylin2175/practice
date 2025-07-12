package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	job   string
	hobby string
	Human
}

func (h Human) justFunc() {
	fmt.Println("Этот метод может быть вызван как от объекта структуры Human, так и от объекта структуры Action")
}

func main() {
	vasya := Human{}
	act := Action{}
	vasya.justFunc()
	act.justFunc()
}
