package main

import "fmt"

func main() {
	var w1 float64
	var w2 float64
	var w3 float64

	fmt.Println("Введите вес основного багажа:")
	fmt.Scan(&w1)

	fmt.Println("Введите вес ручной клади:")
	fmt.Scan(&w2)

	fmt.Println("Введите вес доп. ручной клади:")
	fmt.Scan(&w3)

	n := w1 + w2 + w3
	fmt.Println(n)
}
