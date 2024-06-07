package main

import "fmt"

func main() {
	age := 26
	agePointer := &age // *int

	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age -= 18
}
