package main

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("Personne [Nom:%s, Pays:%s]", p.Name, p.Country)
}

func main() {
	francois := Person{"François", "France"}
	fmt.Println(francois)
}
