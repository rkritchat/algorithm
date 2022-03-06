package main

import (
	"algorithm/tree/node"
)

func main() {
	t := node.NewTree(5)
	t.Add(3)
	t.Add(20)
	t.Add(1)
	t.Add(2)
	t.Add(4)
	t.Print(0)
}
