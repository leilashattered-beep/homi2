package main

import "fmt"

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

func weight(a []InventoryItem) float64 {
	sum := 0.0

	for i := 0; i < len(a); i++ {

		sum = sum + a[i].Weight
	}

	return sum
}

func main() {
	a := []InventoryItem{
		{"меч", 5.5, false},
		{"щит", 8.0, false},
		{"жилетка", 5.0, false},
		{"алмаз", 0.5, true},
		{"амулет", 2.0, true},
	}

	s := weight(a)

	fmt.Println("общий вес:", s)
}
