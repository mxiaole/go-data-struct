package main

import (
	"data-struct/tree"
	"fmt"
)

func main() {
	var t *tree.BinarySearchTree
	root := t.Init()

	fmt.Println(root.GetSize())

	root.Insert(10)
	fmt.Println(root.GetSize())
	root.Insert(20)
	fmt.Println(root.GetSize())
	root.Insert(30)
	fmt.Println(root.GetSize())

	root.PreOrder(root.Root)
}
