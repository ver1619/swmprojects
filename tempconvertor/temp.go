package main

import "fmt"

func main() {
	var celsius float64
	var fahrenheit float64

	var convertionmethod int
	fmt.Print("Choose conversion method (1: Celsius to Fahrenheit, 2: Fahrenheit to Celsius): ")
	fmt.Scanf("%d", &convertionmethod)

	switch convertionmethod {
	case 1:
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scanf("%f", &celsius)
		fahrenheit = (celsius * 9 / 5) + 32
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", celsius, fahrenheit)
	case 2:
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scanf("%f", &fahrenheit)
		celsius = (fahrenheit - 32) * 5 / 9
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", fahrenheit, celsius)
	default:
		fmt.Println("Invalid conversion method selected.")
	}
}
