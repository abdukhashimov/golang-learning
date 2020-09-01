package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func (r rect) allSum() int {
	return r.width + r.height
}

func main() {
	r := rect{height: 2, width: 4}
	fmt.Println(r.allSum())
	fmt.Println("area: ", r.area())
	fmt.Println("perimetr: ", r.perim())
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("Hello World")
}
