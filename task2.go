package tasks

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type Company struct {
	companyName string
	employees   []employee
}

func CreateEmployee(name string, salary int, position string) employee {
	employee1 := employee{name, salary, position}
	return employee1
}

func PrintCompany(c Company) {
	fmt.Println("Company Name:", c.companyName)
	fmt.Println("----------------------------------------------")
	for _, v := range c.employees {
		fmt.Println("Employee Name:", v.name)
		fmt.Println("Employee Salary:", v.salary)
		fmt.Println("Employee Position:", v.position)
		fmt.Println("==============================================")
	}
}

func CreateEmployeeArray() []employee {
	employeeArray := []employee{
		CreateEmployee("Anas", 10000, "Project Manager"),
		CreateEmployee("Shame", 20000, "Frontend Developer"),
		CreateEmployee("Ali", 30000, "Backend Developer"),
	}
	return employeeArray
}
