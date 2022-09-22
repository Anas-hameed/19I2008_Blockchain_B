package tasks

import (
	"fmt"
)

// declare a struct student
type student struct {
	name    string
	age     int
	address string
}

// Write a function that initializes a Student struct and returns it
func InitStudent(name string, age int, address string) student {
	// create a student struct
	student1 := student{name, age, address}
	return student1
}

// pass the student struct to a function that prints the student struct
func PrintStudent(s1 student) {
	fmt.Printf("Name: %s Age: %d Address: %s", s1.name, s1.age, s1.address)
}
