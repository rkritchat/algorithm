package double_node

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type LinkList struct {
	Head *Node
	Len  int
}
