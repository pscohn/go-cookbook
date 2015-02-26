package ds

type iterable interface {
	Iterate()
	Next()
}

// Linked list

type Node struct {
	Data string
	Next *Node
}

// Stack - linked list implementation

type StackLinked struct {
	First *Node
}

func (s *StackLinked) Push(data string) {
	n := Node{data, s.First}
	s.First = &n
}

func (s *StackLinked) Pop() string {
	ret := s.First.Data
	s.First = s.First.Next
	return ret
}

// Stack - array implmentation

type Stack struct {
	// should resize array x2 when full,
	// halve when 1/4 full
	Stack []string
}

func (s *Stack) Push(data string) {
	s.Stack = append(s.Stack, data)
}

func (s *Stack) Pop() string {
	lastindex := len(s.Stack) - 1
	oldlast := s.Stack[lastindex]
	s.Stack = s.Stack[:lastindex]
	return oldlast
}

//Queue - linked list

//type Queue struct {
//	First *Node
//	Last  *Node
//}
//
//func (q *Queue) Enqueue(item string) {
//	node := Node{item, nil}
//	oldLast := q.Last
//	q.Last = &node
//	oldLast.Node.Next = q.Last
//}
//
//func (q *Queue) Dequeue() string {
//
//}
//
//func (q *Queue) isEmpty() bool {
//
//}
