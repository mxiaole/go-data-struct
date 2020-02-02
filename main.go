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
	root.Insert(20)
	root.Insert(30)
	root.Insert(15)
	fmt.Println(root.GetSize())
	root.PreOrder(root.Root)
	fmt.Println(root.IsEmpty())

	fmt.Println(root.Contains(15))
	fmt.Println(root.Contains(34))
}
