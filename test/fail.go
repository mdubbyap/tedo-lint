package test

import (
	"fmt"
)

type test struct {
	i int
}

func (t *test) String() string {
	return fmt.Sprintf("foo %d", t.i)
}

func (t test) Foo() {
	fmt.Println("ho")
}
