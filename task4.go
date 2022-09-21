package tasks

import (
	"crypto/sha256"
	"fmt"
)

type Student_4 struct {
	rollnumber int
	name       string
	address    string
	courses    []string
}

type StudentList_4 struct {
	list []*Student_4
}

// new student
func newStudent_4(rollnumber int, name string, address string, courses []string) *Student_4 {
	s := new(Student_4)
	s.rollnumber = rollnumber
	s.name = name
	s.address = address
	s.courses = courses
	return s
}

// create student
func (ls *StudentList_4) createStudent(rollnumber int, name string, address string, course []string) *Student_4 {
	st := newStudent_4(rollnumber, name, address, course)
	ls.list = append(ls.list, st)
	return st
}

// print the student list in the given format
// print hash of the each student in the list
func printStudentList_4(ls *StudentList_4) {
	for i, v := range ls.list {
		fmt.Println("================================= list", i, "================================")
		fmt.Println("student rollnumber:", v.rollnumber)
		fmt.Println("student name:", v.name)
		fmt.Println("student address:", v.address)
		fmt.Println("student courses:", v.courses)
		// convert courses array to string
		courses := ""
		for _, course := range v.courses {
			courses += course
		}
		// create hash of student name, adress, courses, rollnumber
		hash := sha256.Sum256([]byte(v.name + v.address + string(v.rollnumber) + courses))
		// println("student hash:", hash)
		fmt.Printf("student hash: %x\n", hash)

	}
}
