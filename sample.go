package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
}

func main() {
	// sum(1, 2, 3, 4, 6)
	// nums := []int{1, 2, 3, 4, 5}
	// sum(nums...)
	myOutput := func() int {
		return 1
	}
	fmt.Println(myOutput())
	fmt.Println("Hello World")
}
