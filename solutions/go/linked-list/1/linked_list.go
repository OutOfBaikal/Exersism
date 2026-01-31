package linkedlist

import (
    "errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
    head *Node
    tail *Node
}

type Node struct {
    prev *Node
    next *Node
    Value any
}

func NewList(elements ...any) *List {
    if len(elements) == 0 {
        return &List{}
    }
    
    head := &Node{Value: elements[0]}
    list := &List{head: head}
    current := head
    
	for i := 1; i < len(elements); i++ {
        newNode := &Node {
            Value: elements[i],
            prev: current,
        }
        current.next = newNode
        current = newNode
    }
    list.tail = current
    return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	new_head := &Node{Value: v}
    if l.head == nil {
        l.head = new_head
        l.tail = new_head
    } else {
        new_head.next = l.head
        l.head.prev = new_head
        l.head = new_head
    }
}

func (l *List) Push(v any) {
	new_tail := &Node{Value: v}
    if l.tail == nil {
        l.head = new_tail
        l.tail = new_tail
    } else {
        new_tail.prev = l.tail
        l.tail.next = new_tail
        l.tail = new_tail
    }
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
        return nil, errors.New("No head!")
    }

    old_head := l.head
    value := old_head.Value

    if old_head.next == nil {
        l.head = nil
        l.tail = nil
    } else {
        l.head = old_head.next
        l.head.prev = nil
        old_head.next = nil
    }

    return value, nil
}

func (l *List) Pop() (any, error) {
	if l.tail == nil {
        return nil, errors.New("No tail!")
    }

    old_tail := l.tail
    value := old_tail.Value

    if old_tail.prev == nil {
        l.head = nil
        l.tail = nil
    } else {
        l.tail = old_tail.prev
        l.tail.next = nil
        old_tail.prev = nil 
    }

    return value, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.head.next == nil {
        return
    }

    var prev *Node
    current := l.head

    for current != nil {
        next := current.next
        current.next = prev
        current.prev = next

        prev = current
        current = next
    }

    l.tail = l.head
    l.head = prev
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
