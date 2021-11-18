package main

import (
	"fmt"
	// "mnzgarcia/calculator-kata/calculator"
)

func main() {
	var a, b int
	fmt.Println("Introduce 2 valores:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(Add(a, b))
}
