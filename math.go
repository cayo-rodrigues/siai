package main

import "fmt"

func main() {
	a := 10
	b := 10

	fmt.Printf("%d + %d = %d", a, b, Sum(a, b))
}

func Sum(a int, b int) int {
	return a + b
}
