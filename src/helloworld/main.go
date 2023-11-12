package main

import (
	"fmt"

	"github.com/inherthz/Golang/tree/main/src/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
}
