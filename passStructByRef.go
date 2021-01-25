package main

import (
	"fmt"
)

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {

	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func setAttributes(s *student) {
	s.name = "Ashton Kutcher"
	s.grade = 95
}

func main() {

	var s student
	s.name = "Topher Grace"
	s.grade = 99
	setAttributes(&s)
	printInfo(s)

}
