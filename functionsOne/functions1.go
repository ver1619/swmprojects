package main

import "fmt"

// Swap function to swap two numbers
func Swap(a int, b int) (int, int) {
	return b, a
}

// evenodd function to check if a number is even or odd
func evenodd(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// maxof3 function to find the maximum of three numbers
func maxof3(a int, b int, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

// table function to print multiplication table of a number
func table(n int) {
	for i := 1; i <= 10; i++ {
		result := n * i
		fmt.Printf("%d x %d = %d\n", n, i, result)
	}
}

// sumofn function to calculate the sum of first n natural numbers
func sumofn(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// sumofevenodd function to calculate the sum of even and odd numbers up to n
func sumofevenodd(n int) (int, int) {
	evenSum := 0
	oddSum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}
	return evenSum, oddSum
}

func main() {
	fmt.Println("-----Operations-----")
	fmt.Println("1.Swap Function")
	fmt.Println("2.Even Odd Function")
	fmt.Println("3.Max of 3 Function")
	fmt.Println("4.Multiplication Table Function")
	fmt.Println("5.Sum of N Natural Numbers Function")
	fmt.Println("6.Sum of Even and Odd Numbers Function")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	// Swap Function
	case 1:
		var x int
		var y int
		fmt.Print("Enter first number: ")
		fmt.Scanf("%d", &x)
		fmt.Print("Enter second number: ")
		fmt.Scanf("%d", &y)
		x, y = Swap(x, y)
		fmt.Printf("Before Swapping: %d, %d\n", y, x)
		fmt.Printf("After Swapping: %d, %d\n", x, y)
		// Even Odd Function
	case 2:
		var num int
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &num)
		result := evenodd(num)
		fmt.Printf("The number %d is %s\n", num, result)

		// Max of 3 Function
	case 3:
		var a, b, c int
		fmt.Print("Enter first number: ")
		fmt.Scanf("%d", &a)
		fmt.Print("Enter second number: ")
		fmt.Scanf("%d", &b)
		fmt.Print("Enter third number: ")
		fmt.Scanf("%d", &c)
		max := maxof3(a, b, c)
		fmt.Printf("The maximum of %d, %d and %d is %d\n", a, b, c, max)

		// Multiplication Table Function
	case 4:
		var num int
		fmt.Print("Enter a number to print its multiplication table: ")
		fmt.Scanf("%d", &num)
		table(num)

		// Sum of N Natural Numbers Function
	case 5:
		var num int
		fmt.Print("Enter +ve integer: ")
		fmt.Scanf("%d", &num)
		result := sumofn(num)
		fmt.Printf("The sum of the first %d natural numbers is: %d\n", num, result)

		// Sum of Even and Odd Numbers Function
	case 6:
		var num int
		fmt.Print("Enter a positive integer: ")
		fmt.Scanf("%d", &num)

		if num%2 == 0 {
			evenSum, _ := sumofevenodd(num)
			fmt.Printf("Sum of even numbers from 1 to %d: %d\n", num, evenSum)
		} else {
			_, oddSum := sumofevenodd(num)
			fmt.Printf("Sum of odd numbers from 1 to %d: %d\n", num, oddSum)
		}

	default:
		fmt.Println("Invalid choice")
	}

}
