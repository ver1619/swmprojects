package main

import "fmt"

func main() {
	var name string
	var regNo string
	var class string
	var s1, s2, s3, s4, s5 = "Maths", "Science", "English", "Social Studies", "Computer"
	var m1, m2, m3, m4, m5 int
	var totalMarks int
	var percentage float64
	fmt.Print("Enter Student Name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter Registration Number: ")
	fmt.Scanf("%s", &regNo)
	fmt.Print("Enter Class: ")
	fmt.Scanf("%s", &class)

	fmt.Printf("Enter marks for %s: ", s1)
	fmt.Scanf("%d", &m1)
	fmt.Printf("Enter marks for %s: ", s2)
	fmt.Scanf("%d", &m2)
	fmt.Printf("Enter marks for %s: ", s3)
	fmt.Scanf("%d", &m3)
	fmt.Printf("Enter marks for %s: ", s4)
	fmt.Scanf("%d", &m4)
	fmt.Printf("Enter marks for %s: ", s5)
	fmt.Scanf("%d", &m5)

	totalMarks = m1 + m2 + m3 + m4 + m5
	percentage = (float64(totalMarks) / 500) * 100

	fmt.Println("\n ---- Student Report ---- ")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Registration Number: %s\n", regNo)
	fmt.Printf("Class: %s\n", class)
	fmt.Printf("Total Marks: %d\n", totalMarks)
	fmt.Printf("Percentage: %.2f%%\n", percentage)

	if percentage <= 100 && percentage >= 90 {
		fmt.Println("Grade: A")
	} else if percentage < 90 && percentage >= 80 {
		fmt.Println("Grade: B")
	} else if percentage < 80 && percentage >= 70 {
		fmt.Println("Grade: C")
	} else if percentage < 70 && percentage >= 60 {
		fmt.Println("Grade: D")
	} else if percentage < 60 && percentage >= 50 {
		fmt.Println("Grade: E")
	} else {
		fmt.Println("Grade: F")
	}

	if percentage >= 40 {
		fmt.Println("Result: Pass")
	} else {
		fmt.Println("Result: Fail")
	}
}
