package main

import "fmt"

func votes(v []string) {
	a := 0
	b := 0
	vi := 0

	for i := 0; i < len(v); i++ {
		if v[i] == "Анна" {
			a++
		}
		if v[i] == "Борис" {
			b++
		}
		if v[i] == "Виктор" {
			vi++
		}
	}
	n := len(v)

	fmt.Println("Анна:", a, a*100/n, "%")

	fmt.Println("Борис:", b, b*100/n, "%")

	fmt.Println("Виктор:", vi, vi*100/n, "%")
}

func main() {
	v := []string{"Анна", "Борис", "Анна", "Виктор", "Анна"}
	votes(v)
}
