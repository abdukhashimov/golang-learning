package main

import "fmt"

func main() {
	var str1 = "This is the first string"
	fmt.Println(str1)
	for k, char := range str1 {
		fmt.Printf("%d: %c\n", k, char)
	}
}
