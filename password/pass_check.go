package main

import (
	"fmt"
)

func main() {
	var password string
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	// Loop through each character
	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		} else if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}

	// Calculate score
	score := 0
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}

	// Decide strength
	if len(password) < 8 {
		fmt.Println("Password Strength: Weak (too short)")
	} else if score == 4 {
		fmt.Println("Password Strength: Strong")
	} else if score == 2 || score == 3 {
		fmt.Println("Password Strength: Medium")
	} else {
		fmt.Println("Password Strength: Weak")
	}

	if !hasUpper {
		fmt.Println("Pls add a Uppercase letter")
	}

	if !hasLower {
		fmt.Println("Pls add a Lowercase letter")
	}

	if !hasDigit {
		fmt.Println("Pls add a Digit")
	}

	if !hasSpecial {
		fmt.Println("Pls add a Special character")
	}

}
