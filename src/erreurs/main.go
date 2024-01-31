package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	LastName  string
	FirstName string
}

var ErrNotFound = errors.New("Employé non trouvé")

func main() {
	employee, err := GetEmployee(1)
	// if err != nil {
	// 	fmt.Println("Erreur")
	if errors.Is(err, ErrNotFound) {
		fmt.Println(err)
	} else {
		fmt.Println(*employee)
	}
}

func GetEmployee(id int) (*Employee, error) {
	employee, err := apiGetEmployee(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return employee, err
}

func apiGetEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "foo", FirstName: "bar", ID: 1}
	return &employee, ErrNotFound
}
