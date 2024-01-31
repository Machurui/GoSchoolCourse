package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstName string = "François"
	lastName := "Serra"
	fmt.Println(firstName, lastName)

	const HTTPOk = 200

	// Types

	// int8, int16, int32, int64 (int ou uint <- unsigned)

	var integer16 int16 = 2458

	var integer32 int32 = 12454

	fmt.Println(integer16 + int16(integer32))

	// float32 float64

	const pi = 3.14957
	const avogadro = 6.022e23

	// bool

	const faux bool = false

	// Séquences d'échappement \n \r \t \' \" \\

	var defaultint int
	var defautfloat32 float32
	var defautfloat64 float64
	var defaultbool bool
	var defaultstring string

	fmt.Println(defaultint, defautfloat32, defautfloat64, defaultbool, defaultstring)

	// Cast explicite
	num := "42"
	value, _ := strconv.Atoi(num)

	s := strconv.Itoa(value)

	fmt.Println(num, value, s)
}
