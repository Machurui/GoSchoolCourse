package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	newFile, error := os.Create("go.txt")
	if error != nil {
		fmt.Println("Impossible de créer le fichier")
		return
	}

	defer newFile.Close()

	if _, error = io.WriteString(newFile, "C'est un super test"); error != nil {
		fmt.Println("Impossible d'écrire dans le fichier")
		return
	}

	newFile.Sync()

	for i := 1; i <= 4; i++ {
		defer fmt.Println("defer", i)
		fmt.Println("Standard", i)
	}

	fmt.Println(foo())

	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("Main recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Fin")

}

func foo() (i int) {
	defer func() { i++ }()
	return 1
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic")
		panic("Low supérieur à High")
	}

	defer fmt.Println("defer", high, low)

	fmt.Println("standard", high, low)

	highlow(high, low+1)
}
