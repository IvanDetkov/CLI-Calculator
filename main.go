package main

import "fmt"

func main() {

	for {
		var (
			a      float64
			b      float64
			c      string
			result float64
			choice int
		)

		fmt.Print("Enter the first number: ")
		a = getNumber()

		fmt.Print("Enter the second number: ")
		b = getNumber()

		fmt.Print("Enter the operator (+, -, *, /): ")
		c = getMark()

		// calculations
		switch c {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "/":
			if b == 0 {
				fmt.Println("Error! Division by zero is impossible!")
				return
			}
			result = a / b
		case "*":
			result = a * b
		default:
			fmt.Println("Error! Invalid mark!")
		}
		fmt.Printf("Answer: %.2f %v %.2f = %.2f\n", a, c, b, result)

		// Prompt the user to continue or exit.
		// If the input is invalid, display an error message and repeat the loop.
		fmt.Println("Do you want to continue? 1 - Yes, 2 - No")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			continue
		case 2:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Error! Invalid number!")
		}

	}
}

// getNumber prompts the user to enter a number and validates the input.
// It keeps asking until a valid numeric value is provided.
func getNumber() float64 {
	var num float64
	for {
		_, err := fmt.Scan(&num)
		if err == nil {
			return num
		}
		fmt.Println("Error! Enter a number, not text!")
	}
}

// getMark prompts the user to enter an operator symbol (+, -, *, /).
// It keeps asking until a valid string input is provided.
func getMark() string {
	var mark string
	for {
		_, err := fmt.Scan(&mark)
		if err == nil {
			return mark
		}
		fmt.Println("Error: please enter a symbol!")
	}
}
