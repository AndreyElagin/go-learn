package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) String() string {
	return fmt.Sprintf("{ width: %d, height: %d }", r.width, r.height)
}

func (r *rect) makeWider(times int) {
	r.width = r.width * times
}

func (r *rect) makeHigher(times int) {
	r.height = r.height * times
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
