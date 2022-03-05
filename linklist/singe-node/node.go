package main

import "fmt"

const (
	PRE  = "PREVIOUS"
	CURR = "CURRENT"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	Head *Node
	Len  int
}

func InitNode(val interface{}) *LinkList {
	return &LinkList{
		Head: &Node{
			Data: val,
			Next: nil,
		},
		Len: 1,
	}
}

func (l *LinkList) AddLast(val interface{}) {
	h := l.Head
	for h != nil {
		if h.Next == nil {
			h.Next = &Node{
				Data: val,
				Next: nil,
			}
			l.Len += 1
			break
		}
		h = h.Next
	}
}

func (l *LinkList) AddFirst(val interface{}) {
	l.Head = &Node{
		Data: val,
		Next: l.Head,
	}
	l.Len += 1
}

func (l *LinkList) AddAfterTarget(target interface{}, val interface{}) {
	current := l.Search(target, CURR)
	if current == nil { //case target not found
		l.AddLast(val)
		return
	}
	tmp := current.Next
	current.Next = &Node{
		Data: val,
		Next: tmp,
	}
	l.Len += 1
}

func (l *LinkList) AddBeforeTarget(target interface{}, val interface{}) {
	previous := l.Search(target, PRE)
	if previous == nil { //case target not found
		if l.Head.Data == target {
			l.AddFirst(val)
			return
		}
		l.AddLast(val)
	}
	tmp := previous.Next
	previous.Next = &Node{
		Data: val,
		Next: tmp,
	}
	l.Len += 1
}

func (l *LinkList) Show() {
	h := l.Head
	for h != nil {
		fmt.Println(h.Data)
		h = h.Next
	}
	fmt.Printf("total %v\n", l.Len)
}

func (l *LinkList) Search(target interface{}, on string) *Node {
	h := l.Head
	count := 0
	var previous *Node
	for h != nil {
		if h.Data == target {
			switch on {
			case PRE:
				return previous
			case CURR:
				return h
			default:
				return h
			}
		}
		count += 1
		previous = h
		h = h.Next
	}
	return nil
}

func (l *LinkList) Remove(target interface{}) {

}

func (l *LinkList) Reverse(last *Node, first *Node) {
	if last == nil {
		h := l.Head
		//find last node
		for h != nil {
			if h.Next == nil {
				last = h
				first = h
			}
			h = h.Next
		}
	}

	//search previous
	previous := l.Search(last.Data, PRE)
	if previous == nil {
		l.Head = first
		return
	}

	last.Next = previous
	previous.Next = nil
	l.Reverse(previous, first)
}

func main() {
	n := InitNode(1)
	n.AddBeforeTarget(1, 2)
	n.Show()
}
