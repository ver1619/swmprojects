package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter capacity: ")
	fmt.Scan(&n)
	arr := make([]int, 0, n)

	fmt.Println("Array of capacity", n, "created:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Println(arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Sum:", sum)
	average := float64(sum) / float64(len(arr))
	fmt.Printf("Average: %.2f\n", average)
}
