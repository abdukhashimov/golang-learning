package main

import "fmt"

type Name struct {
	name    string
	age     int
	isAdmin bool
}

type User struct {
	name     string
	age      int
	isAdmin  bool
	lastname string
}

type User2 struct {
	firstname string
	lastname  string
	age       int
	isAdmin   bool
	paidTill  string
}

type User3 struct {
	firstname string
	lastname  string
	age       int
	isAdmin   bool
	paidTill  string
}

type User4 struct {
	firstname string
	lastname  string
	age       int
	isAdmin   bool
	paidTill  string
}

type Name struct {
	field  int
	field2 string
	field3 bool
}

func main() {
	fmt.Println("Hello World")
}
