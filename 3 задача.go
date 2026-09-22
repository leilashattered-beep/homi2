package main

import "fmt"

type Order struct {
	id          int
	items       []int
	total       float64
	address     string
	isCompleted bool
}

func main() {
	orders := make(map[int]Order)
	var id int
	var itemCount int
	var total float64
	var address string
	var isCompleted bool

	fmt.Println("Введите ID заказа:")
	fmt.Scan(&id)

	fmt.Println("Введите количество товаров:")
	fmt.Scan(&itemCount)

	items := make([]int, itemCount)
	for i := 0; i < itemCount; i++ {
		fmt.Printf("Введите ID товара %d:\n", i+1)
		fmt.Scan(&items[i])
	}

	fmt.Println("Введите сумму:")
	fmt.Scan(&total)

	fmt.Println("Введите адрес доставки:")
	fmt.Scan(&address)

	fmt.Println("Заказ готов? (0 = нет, 1 = да):")
	var completed int
	fmt.Scan(&completed)
	isCompleted = completed == 1

	orders[id] = Order{
		id:          id,
		items:       items,
		total:       total,
		address:     address,
		isCompleted: isCompleted,
	}

	fmt.Println("\nВаш заказ:")
	fmt.Println("ID:", orders[id].id)
	fmt.Println("Товары:", orders[id].items)
	fmt.Println("Сумма:", orders[id].total)
	fmt.Println("Адрес:", orders[id].address)
	fmt.Println("Готов:", orders[id].isCompleted)
}
