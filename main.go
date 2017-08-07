package main

import (
	"fmt"

	"github.com/jeffprestes/vendorerradoa/calc"
)

func main() {
	a := 3
	b := 1
	c := 1
	fmt.Println("Quero o Delta também: ", calc.Delta(a, b, c))
}
