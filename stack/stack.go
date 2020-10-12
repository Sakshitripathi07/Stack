package stack

import (
	"log"
)

//MAX defines the maximum size of stack
const MAX = 10

//Stack is of type int
type Stack []int

//Push function pushes an element to the stack
func (s *Stack) Push(i ...int) (int, bool) {
	if len(*s) == MAX {
		log.Println("not pushed")
		return -1, false
	}
	*s = append(*s, i...)
	log.Println("pushed")
	log.Println("stack is ", s, len(*s))
	return len(*s), true
}

//Pop function pops an element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		log.Println("not popped")
		return -1, false
	}
	top := len(*s) - 1
	popped := (*s)[top]
	*s = (*s)[0:top]
	log.Println("popped", popped)
	return len(*s), true
}
