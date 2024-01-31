package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 27
	if x%2 == 0 {
		fmt.Println("X est pair.")
	} else if x == 0 {
		fmt.Println("x est égal à 0.")
	} else {
		fmt.Println("x est impair.")
	}

	if num := -1; num < 0 {
		fmt.Println(num, "neg")
	} else if num > 0 {
		fmt.Println(num, "pos")
	} else {
		fmt.Println(num, "zero")
	}

	// Switch
	sec := time.Now().Unix()

	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("un")
	default:
		fmt.Println("autre")
	}

	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday":
		fmt.Println("Semaine")
	default:
		fmt.Println("Week-end")
	}

	switch num := 15; {
	case num < 50:
		fmt.Println("Inférieur à 50")
		fallthrough
	case num > 100:
		fmt.Println("Supérieur à 100")

	}

	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println(sum)

	// While
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(10)
		fmt.Println(num)
	}

	// While true
	for {
		if num = rand.Int63n(10); num == 5 {
			fmt.Println("Terminé")
			break
		}
	}

	fmt.Println(num)

	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}

		sum += i
	}
	fmt.Println(sum)
}
