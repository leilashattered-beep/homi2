package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func salary(employees []Employee) (float64, float64) {
	sum := 0.0

	for i := 0; i < len(employees); i++ {

		sum = sum + employees[i].Salary
	}

	average := sum / float64(len(employees))
	return sum, average
}

func main() {
	employees := []Employee{
		{1, "Елена", "Повар", 60000},
		{2, "Паша", "Менеджер", 50000},
		{3, "Маша", "Официант", 55000},
	}

	sum, average := salary(employees)

	fmt.Println("Общий фонд оплаты труда:", sum)
	fmt.Println("Средняя зарплата:", average)
}
