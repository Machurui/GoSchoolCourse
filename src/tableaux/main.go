package main

import (
	"fmt"
)

func main() {
	var tab [3]int
	tab[1] = 10
	fmt.Println(tab[0])
	fmt.Println(tab[1])
	fmt.Println(tab[len(tab)-1])

	cities := [...]string{"Paris", "Tokyo"}
	fmt.Println("Villes", cities)

	numbers := [...]int{99: -1}
	fmt.Println(numbers[0], numbers[99], len(numbers))

	genres := []string{"Homme", "Femme", "Autres"}
	fmt.Println(genres, len(genres), cap(genres))

	x := [3]int{1, 2, 3}
	s := x[:]
	fmt.Println(s)

	months := []string{"Ja", "Fe", "Ma", "Av", "Mai", "Ju", "Jl", "Ao", "Se", "Oc", "No", "De"}
	trim1 := months[0:3]
	trim2 := months[3:6]
	trim3 := months[6:9]
	trim4 := months[9:12]
	fmt.Println(trim1, trim2, trim3, trim4)
	fmt.Println(len(trim1), len(trim2), len(trim3), len(trim4))
	fmt.Println(cap(trim1), cap(trim2), cap(trim3), cap(trim4))

	trim2ex := trim2[:4]
	fmt.Println(trim2ex, len(trim2ex), cap(trim2ex))

	var nums []int
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(nums), nums)
	}

	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	if remove < len(letters) {
		fmt.Println("Avant", letters, "Suppr", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("Après", letters)
	}

	slice1 := letters[:2]
	slice2 := letters[1:4]
	fmt.Println(letters, slice1, slice2)
	slice1[1] = "P"
	fmt.Println(letters, slice1, slice2)

	slice3 := make([]string, 3)
	copy(slice3, letters[1:4])
	fmt.Println(letters, slice3)
	slice3[1] = "S"
	fmt.Println(letters, slice3)
}
