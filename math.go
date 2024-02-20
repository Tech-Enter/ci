package main

import "fmt"

func main() {
	fmt.Println(Soma(2, 3))
	// Output: 5
	// go test
	// go test -v
}

func Soma(a int, b int) int {
	return a + b
}
