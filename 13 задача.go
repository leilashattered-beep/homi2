package main

import "fmt"

func main() {
	expenses := map[string]float64{
		"еда":         15000,
		"транспорт":   5000,
		"развлечения": 3000,
	}

	expenses["еда"] = expenses["еда"] + 2000

	sum := 0.0

	for category, money := range expenses {
		fmt.Println(category, money)
		sum = sum + money
	}

	fmt.Println("итого:", sum)
}
