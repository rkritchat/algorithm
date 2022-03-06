package node

import "fmt"

type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func NewTree(val int) *Node {
	return &Node{
		Value: val,
	}
}

func (n *Node) Print(sequence int) {
	fmt.Printf("[%v]%v\n", sequence, n.Value)
	if n.Left != nil {
		fmt.Println("go left")
		n.Left.Print(sequence + 1)
	}
	if n.Right != nil {
		fmt.Println("go right")
		n.Right.Print(sequence + 1)
	}
}

func (n *Node) Add(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
			return
		}
		n.Left.Add(val)
		return
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: val}
			return
		}
		n.Right.Add(val)
		return
	}
	//equald, do nothing
}
