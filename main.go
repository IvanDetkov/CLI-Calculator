package main

import (
	"fmt"
)

func main() {

	for {
		var (
			a      float64
			b      float64
			c      string
			result float64
			choice int
		)

		err := getInput("Enter the first number: ", &a)
		if err != nil {
			fmt.Println("Input error!", err)
			continue
		}

		err = getInput("Enter the second number: ", &b)
		if err != nil {
			fmt.Println("Input error!", err)
			continue
		}

		err = getInput("Enter the operator (+, -, *, /): ", &c)
		if err != nil {
			fmt.Println("Input error!", err)
			continue
		}

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

// getInput prompts the user for input from [io.Stdin].
func getInput(msg string, inp any) error {
	for {
		fmt.Print(msg)

		_, err := fmt.Scan(inp)

		return err
	}
}
