package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numberOne := os.Args[1]
	numberTwo := os.Args[2]

	fmt.Println(numberOne, numberTwo, sum(numberOne, numberTwo))

	sum, mul := calc(numberOne, numberTwo)
	fmt.Println(sum, mul)

	name := "François"
	update(name)
	fmt.Println(name)

	updatePointer((&name))
	fmt.Println(name)
}

func sum(numberOne string, numberTwo string) (result int) {
	intOne, _ := strconv.Atoi(numberOne)
	intTwo, _ := strconv.Atoi(numberTwo)

	result = intOne + intTwo
	return
}

func calc(numberOne string, numberTwo string) (sum int, mul int) {
	intOne, _ := strconv.Atoi(numberOne)
	intTwo, _ := strconv.Atoi(numberTwo)

	sum = intOne + intTwo
	mul = intOne * intTwo
	return
}

func update(name string) {
	name = "David"
}

func updatePointer(name *string) {
	*name = "David"
}
