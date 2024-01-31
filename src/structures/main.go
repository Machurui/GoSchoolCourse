package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID        int
	LastName  string
	FirstName string `json:"Prénom"`
}

type Responsible struct {
	Employee
	ServiceName string
}

func main() {
	francois := Employee{1, "Serra", "François"}
	fmt.Println(francois)
	fmt.Println(francois.FirstName)

	thomas := Employee{LastName: "Calmettes", FirstName: "Thomas"}
	fmt.Println(thomas)

	thomas.ID = 2
	fmt.Println(thomas.ID)

	var nicolas Responsible
	nicolas.FirstName = "Nicolas"
	fmt.Println(nicolas)

	// Json
	people := []Responsible{
		{
			Employee:    Employee{ID: 1, LastName: "Gros", FirstName: "Christian"},
			ServiceName: "Info",
		},
		{
			Employee:    Employee{ID: 1, LastName: "Mebani", FirstName: "Samir"},
			ServiceName: "RH",
		},
	}

	data, _ := json.Marshal(people)

	fmt.Printf("%s\n", data)

	var decode []Employee

	json.Unmarshal(data, &decode)
	fmt.Printf("%v", decode)
}
