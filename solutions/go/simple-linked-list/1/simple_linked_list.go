package linkedlist

import (
    "errors"
)

// Define the List and Element types here.
type List struct {
    head *Element
}

type Element struct {
    next *Element
    Value int
}

func New(elements []int) *List {
    if len(elements) == 0 {
        return &List{}
    }
    
	head := &Element{Value: elements[0]}
    list := &List{head: head}
    current := head

    for i := 1; i < len(elements); i++ {
        newNode := &Element {
            Value: elements[i],
        }
        current.next = newNode
        current = newNode
    }
    return list
}

func (l *List) Size() int {
    if l == nil || l.head == nil{
        return 0
    }
    res := 1
    buf := l.head
    for buf.next != nil {
        buf = buf.next
        res += 1
    }

    return res
}

func (l *List) Push(element int) {
	newElement := &Element{Value: element}

    if l.head == nil {
        l.head = newElement
        return
    }

    buf := l.head
    for buf.next != nil {
        buf = buf.next
    }
    buf.next = newElement
}

func (l *List) Pop() (int, error) {
	if l == nil || l.head == nil {
        return 0, errors.New("List is empty!")
    }

    if l.head.next == nil {
        value := l.head.Value
        l.head = nil 
        return value, nil
    }

    buf := l.head
    for buf.next.next != nil {
        buf = buf.next
    }
    value := buf.next.Value
    buf.next = nil
    return value, nil
}

func (l *List) Array() []int {
	var array []int
    buf := l.head
    for buf != nil {
        array = append(array, buf.Value)
        buf = buf.next
    }

    return array
}

func (l *List) Reverse() *List {
    if l == nil || l.head == nil {
        return l
    }

    var prev *Element
    current := l.head
    for current != nil {
        next := current.next 
        current.next = prev  
        prev = current       
        current = next       
    }
    return &List{head: prev}
}
