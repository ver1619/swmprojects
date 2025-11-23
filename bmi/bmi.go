package main

import "fmt"

func calcBMI(weight, height float64) float64 {
	fmt.Println("Weight:", weight, "(kg)")
	fmt.Println("Height:", height, "(m)")
	return weight / (height * height)

}

func category(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi < 24.9 {
		return "Normal"
	} else if bmi < 29.9 {
		return "Overweight"
	}
	return "Obese"
}

func main() {
	var weight float64
	var height float64
	fmt.Println("-----Instructions-----")
	fmt.Println("1.Please make sure you have provided Height in meters(m)")
	fmt.Println("2.Please make sure you have provided Weight in Kgs(kg))")
	fmt.Print("Enter Weight:")
	fmt.Scan(&weight)
	fmt.Print("Enter Height:")
	fmt.Scan(&height)
	bmi := calcBMI(weight, height)
	cat := category(bmi)
	fmt.Println("BMI:", bmi)
	fmt.Println("Category:", cat)
}
