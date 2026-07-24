package main

import "fmt"

func main() {
	// ввод первого числа с проверкой
	var num1 float64
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка! Введите число, а не текст!")
		return
	}

	// ввод второго числа с проверкой
	var num2 float64
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка! Введите число, а не текст!")
		return
	}

	// ввод знака с проверкой
	var mark string
	fmt.Print("Введите знак (+, -, *, /): ")
	fmt.Scan(&mark)
	if mark != "+" && mark != "-" && mark != "*" && mark != "/" {
		fmt.Printf("Ошибка! Неверный знак: %v\n", mark)
	}

	// вычесления
	var result float64
	switch mark {
	case "+":
		result = num1 + num2
		fmt.Printf("Ответ: %.2f + %.2f = %.2f\n", num1, num2, result)
	case "-":
		result = num1 - num2
		fmt.Printf("Ответ: %.2f - %.2f = %.2f\n", num1, num2, result)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка! Деление на ноль невозможно!")
			return
		}
		result = num1 / num2
		fmt.Printf("Ответ: %.2f / %.2f = %.2f\n", num1, num2, result)
	case "*":
		result = num1 * num2
		fmt.Printf("Ответ: %.2f * %.2f = %.2f\n", num1, num2, result)
	}
}
