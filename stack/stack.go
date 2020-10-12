package stack

import (
	"log"
)

//MAX defines the maximum size of stack
const MAX = 10

//Stack is of type int
type Stack []int

//Push function pushes an element to the stack
func (s *Stack) Push(i int) (string, bool) {
	if len(*s) == MAX {
		log.Println("not pushed")
		return "stack is full", false
	}
	*s = append(*s, i)
	log.Println("pushed", i)
	return "pushed", true
}

//Pop function pops an element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		log.Println("not popped")
		return 0, false
	}
	top := len(*s) - 1
	popped := (*s)[top]
	*s = (*s)[0:top]
	log.Println("popped", popped)
	return popped, true
}
