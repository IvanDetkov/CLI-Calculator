package main

import "fmt"

func main() {
	// input of the first number with validation
	var num1 float64
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error! Enter a number, not text!")
		return
	}

	// input of the second number with validation
	var num2 float64
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error! Enter a number, not text!")
		return
	}

	// input of the operator with validation
	var mark string
	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scan(&mark)
	if mark != "+" && mark != "-" && mark != "*" && mark != "/" {
		fmt.Printf("Error! Invalid operator: %v\n", mark)
	}

	// calculations
	var result float64
	switch mark {
	case "+":
		result = num1 + num2
		fmt.Printf("Answer: %.2f + %.2f = %.2f\n", num1, num2, result)
	case "-":
		result = num1 - num2
		fmt.Printf("Answer: %.2f - %.2f = %.2f\n", num1, num2, result)
	case "/":
		if num2 == 0 {
			fmt.Println("Error! Division by zero is impossible!")
			return
		}
		result = num1 / num2
		fmt.Printf("Answer: %.2f / %.2f = %.2f\n", num1, num2, result)
	case "*":
		result = num1 * num2
		fmt.Printf("Answer: %.2f * %.2f = %.2f\n", num1, num2, result)
	}
}
