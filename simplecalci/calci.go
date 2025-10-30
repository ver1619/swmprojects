//Simple Calculator(+,-,*,/,%)

package main

import "fmt"

func main() {
	var a, b int
	var o string
	var m string

	fmt.Print("Enter first number:")
	fmt.Scan(&a)
	fmt.Print("Enter second number:")
	fmt.Scan(&b)
	fmt.Print("Enter operator (+,-,*,/,%):")
	fmt.Scan(&o)

	fmt.Print("Choose method (switch/if-else):")
	fmt.Scan(&m)
	//Switch Case Method

	if m == "switch" {
		switch o {
		case "+":
			fmt.Printf("Result : %d + %d = %d\n", a, b, a+b)
		case "-":
			fmt.Printf("Result : %d - %d = %d\n", a, b, a-b)
		case "*":
			fmt.Printf("Result : %d * %d = %d\n", a, b, a*b)
		case "/":
			if b != 0 {
				fmt.Printf("Result : %d / %d = %.2f\n", a, b, float32(a)/float32(b))
			} else {
				fmt.Println("Error: Division by zero")
			}
		case "%":
			if b != 0 {
				fmt.Printf("Result : %d %% %d = %.2f\n", a, b, float32(a%b))
			} else {
				fmt.Println("Error: Division by zero")
			}
		default:
			fmt.Println("Invalid operator")
		}

		//If-Else Method

	} else if m == "if-else" {
		if o == "+" {
			fmt.Printf("Result : %d + %d = %d\n", a, b, a+b)
		} else if o == "-" {
			fmt.Printf("Result : %d - %d = %d\n", a, b, a-b)
		} else if o == "*" {
			fmt.Printf("Result : %d * %d = %d\n", a, b, a*b)
		} else if o == "/" {
			if b != 0 {
				fmt.Printf("Result : %d / %d = %.2f\n", a, b, float32(a)/float32(b))
			} else {
				fmt.Println("Error: Division by zero")
			}
		} else if o == "%" {
			if b != 0 {
				fmt.Printf("Result : %d %% %d = %.2f\n", a, b, float32(a%b))
			} else {
				fmt.Println("Error: Division by zero")
			}
		} else {
			fmt.Println("Invalid operator")
		}
	} else {
		fmt.Println("Invalid method choice")
	}
}
