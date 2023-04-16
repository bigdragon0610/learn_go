package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

type StringerError struct {
	v interface{}
}

func (e StringerError) Error() string {
	return fmt.Sprintf("%TはStringerに変換できません", e.v)
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, StringerError{v}
}

type Func func() string

func (f Func) String() string {
	return f()
}

func main() {
	if s, err := ToStringer(Func(func() string {
		return "a"
	})); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(s)
	}
	if s, err := ToStringer(func() string {
		return "a"
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(s)
	}
}
