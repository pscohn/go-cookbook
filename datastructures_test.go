package ds

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	a := Node{"Hello", nil}
	b := Node{"There", &a}
	fmt.Println(&a)
	fmt.Println(&a.Data)
	fmt.Println(&b.Next.Data)

	if &a != b.Next {
		t.Error("b.next does not equal &")
	}
}

func TestStackLinked(t *testing.T) {
	a := Node{"Hello", nil}
	b := Node{"There", &a}
	c := Node{"My", &b}
	d := Node{"Friend", &c}
	s := Stack{&d}
	fmt.Println(s)
}
