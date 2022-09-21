package tasks

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func createEmployee(name string, salary int, position string) employee {
	employee1 := employee{name, salary, position}
	return employee1
}

func printCompany(c company) {
	fmt.Println("Company Name:", c.companyName)
	fmt.Println("----------------------------------------------")
	for _, v := range c.employees {
		fmt.Println("Employee Name:", v.name)
		fmt.Println("Employee Salary:", v.salary)
		fmt.Println("Employee Position:", v.position)
		fmt.Println("==============================================")
	}
}

func createEmployeeArray() []employee {
	employeeArray := []employee{
		createEmployee("Anas", 10000, "Project Manager"),
		createEmployee("Shame", 20000, "Frontend Developer"),
		createEmployee("Ali", 30000, "Backend Developer"),
	}
	return employeeArray
}
