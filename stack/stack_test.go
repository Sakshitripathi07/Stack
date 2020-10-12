package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	var s Stack
	for _, table := range []struct {
		value []int
	}{
		{value: []int{}},
		{value: []int{1}},
		{value: []int{1, 2, 3, 4, 5}},
		{value: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{value: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	} {
		if size, _ := s.Push(table.value...); size >= MAX {
			t.Fatalf("size of stack exceeded maximum stack size %d ",MAX)
		}
	}
}

func TestPop(t *testing.T){
	var s Stack
	if len,_:=s.Pop();len<=0{
		t.Fatalf("Stack is empty nothing to pop")
	}
}