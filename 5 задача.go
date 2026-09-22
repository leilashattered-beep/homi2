package main

import (
	"fmt"
	"strings"
)

func validateUser(name string, age int, email string) error {
	if name == "" {
		return fmt.Errorf("имя пустое")
	}
	if len(name) >= 50 {
		return fmt.Errorf("имя слишком длинное")
	}
	if age < 18 || age > 90 {
		return fmt.Errorf("неправильный возраст")
	}
	if strings.Index(email, "@") == -1 {
		return fmt.Errorf("неправильная почта")
	}
	return nil
}

func main() {
	err := validateUser("мария", 19, "mariha@mail.ru")

	if err != nil {
		fmt.Println(err)
	} 
	else {
		fmt.Println("данные правильные")
	}
}