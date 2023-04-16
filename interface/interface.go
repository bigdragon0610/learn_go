package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (i I) String() string {
	return "I"
}

type B bool

func (b B) String() string {
	return "B"
}

type S string

func (s S) String() string {
	return "S"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case B:
		fmt.Println(bool(v), "B")
	case S:
		fmt.Println(string(v), "S")
	}
}

func main() {
	F(I(100))
	F(B(true))
	F(S("hoge"))
}
