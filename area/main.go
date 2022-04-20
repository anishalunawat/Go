package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	b float64
	h float64
}

type square struct {
	s float64
}

func main() {
	tr := triangle{b: 10, h: 10}
	sq := square{s: 10}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.b * tr.h
}
func (sq square) getArea() float64 {
	return sq.s * sq.s
}
