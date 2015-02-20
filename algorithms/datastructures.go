package ds

type Node struct {
	Data string
	Next *Node
}

type StackLinked struct {
	First *Node
}

func (s *StackLinked) Push(data string) {
	n := Node{data, s.First}
	s.First = &n
}
