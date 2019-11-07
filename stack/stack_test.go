package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	sk := NewStack()
	sk.Push(1)
	sk.Push(100)
	sk.Push(20)
	sk.Pop()
	fmt.Println(sk.Top(), sk.IsEmpty(), sk)
}
