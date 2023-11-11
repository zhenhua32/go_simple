package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	salary float64
}

func main() {
	emp := Employee{
		Person: Person{
			name: "John Doe",
			age:  30,
		},
		salary: 5000.0,
	}

	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Salary:", emp.salary)
}
