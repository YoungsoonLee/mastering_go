package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size++
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp // !!!
	size++
	return true
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		result := t.Value
		t = nil
		return result, true
	}

	result := t.Value
	t = t.Next
	size--
	return result, true
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("Empty stack")
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, b := Pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop failed")
	}

	Push(100)
	traverse(stack)
}
