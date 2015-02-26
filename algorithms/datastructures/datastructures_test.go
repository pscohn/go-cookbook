package ds

import "testing"

func TestNode(t *testing.T) {
	a := Node{"Hello", nil}
	b := Node{"There", &a}
	//	fmt.Println(&a)
	//	fmt.Println(&a.Data)
	//	fmt.Println(&b.Next.Data)

	if &a != b.Next {
		t.Error("b.next does not equal &")
	}
}

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push("Hohoho")

	if s.Pop() != "Hohoho" {
		t.Error("Stack.Pop() does not return last pushed entry")
	}

}

func TestStackLinked(t *testing.T) {
	a := Node{"Hello", nil}
	b := Node{"There", &a}
	c := Node{"My", &b}
	d := Node{"Friend", &c}
	s := StackLinked{&d}

	if s.First != &d {
		t.Error("stack.First does not equal &Node")
	}

	s.Push("Hohoho")
	if s.First.Data != "Hohoho" {
		t.Error("stack.First.Data does not match s.Push(Data)")
	}

	if s.Pop() != "Hohoho" {
		t.Error("stack.Pop() does not return last pushed entry")
	}

}
