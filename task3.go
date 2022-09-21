package tasks

type Student struct {
	rollnumber int
	name       string
	address    string
}

type StudentList struct {
	list []*Student
}

// new student
func newStudent(rollnumber int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollnumber
	s.name = name
	s.address = address
	return s
}

// create student
func (ls *StudentList) createStudent(rollnumber int, name string, address string) *Student {
	st := newStudent(rollnumber, name, address)
	ls.list = append(ls.list, st)
	return st
}

// print the student list in the given format
func printStudentList(ls *StudentList) {
	for i, v := range ls.list {
		println("================================= list", i, "================================")
		println("student rollnumber:", v.rollnumber)
		println("student name:", v.name)
		println("student address:", v.address)
	}
}
