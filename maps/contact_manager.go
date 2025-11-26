package main

import "fmt"

func add_contact(n map[string]int) {
	var name string
	var number int
	fmt.Print("Enter Name:")
	fmt.Scan(&name)
	fmt.Print("Enter Number:")
	fmt.Scan(&number)

	_, exists := n[name]
	if exists {
		fmt.Println("Contact already exists. Use update option.")
		return
	}
	n[name] = number
	fmt.Println("New Contact Added")
	for key, val := range n {
		fmt.Println("Name:", key, "->", val)
	}
}

func delete_contact(n map[string]int, name string) {
	_, exists := n[name]
	if !exists {
		fmt.Println("Contact does not exists.")
		return
	}

	delete(n, name)
	fmt.Println("Contact Deleted:", name)
}

func update_contact(n map[string]int, name string) {
	_, exists := n[name]
	if !exists {
		fmt.Println("Cannot Update. Contact Does Not Exits!")
		return
	}
	var new_number int
	fmt.Print("Enter Number:")
	fmt.Scan(&new_number)
	n[name] = new_number
	fmt.Println("Updated Contact")
	fmt.Println("Contact Details:")
	fmt.Println("Name:", name, "->", n[name])

}

func search_contact(m map[string]int, name string) {
	val, exists := m[name]
	if !exists {
		fmt.Println("Contact not found:", name)
		return
	}
	fmt.Println("Found:", name, "->", val)
}

func main() {
	a := map[string]int{
		"Vinay": 234929173,
	}
	fmt.Println("Welcome! To your Personal Phone Manager")
	for {
		fmt.Println("===Operations===")
		fmt.Println("1.Add Contact")
		fmt.Println("2.Delete Contact")
		fmt.Println("3.Update Contact")
		fmt.Println("4.Search Contact")
		fmt.Println("5.List Contacts")
		fmt.Println("6.Exit")

		var menu int
		fmt.Println("Enter the operation you need to perform:")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			add_contact(a)

		case 2:
			var k string
			fmt.Print("Enter Contact to delete:")
			fmt.Scan(&k)
			delete_contact(a, k)

		case 3:
			var k string
			fmt.Print("Enter Contact name to update:")
			fmt.Scan(&k)
			update_contact(a, k)

		case 4:
			var k string
			fmt.Print("Enter Contact name to search:")
			fmt.Scan(&k)
			search_contact(a, k)

		case 5:
			for key, val := range a {
				fmt.Println("Name:", key, "->", val)
			}

		case 6:
			fmt.Println("Exiting...")
			return
		}
	}
}
