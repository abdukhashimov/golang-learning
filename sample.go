package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(5))
	fmt.Println("Hello World")
}
