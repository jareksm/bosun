package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type A struct {
	F1 int
	F2 int
	d  string
}

type B struct {
	A
	F3, F4 int
	F5     map[int]A
	c      int
}

type C struct {
	B
	c string
}

func main() {
	a := B{
		A: A{
			F1: 1,
			F2: 2,
		},
		F3: 3,
		F4: 4,
		F5: map[int]A{
			0: A{
				F1: 11,
				F2: 12,
			},
			4: A{
				F1: 41,
				F2: 42,
			},
		},
		c: 99,
	}
	a.d = "jarek"
	b := C{}
	fmt.Println(a.F1, a.F2, a.A.F2, a.F3, a.F5[4].F1, a.c, a.d)
	spew.Dump(a)
	spew.Dump(b)
	fmt.Println(b.c, b.F1, b.B.c, b.F3, b.d)
}
