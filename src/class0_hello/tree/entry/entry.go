package main

import (
	"fmt"
	"tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //new()新建对象
	//无论是地址还是结构本身,都用.来访问成员

	root.Left.Right = tree.CreateNode(1)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.Print()
	root.SetValue(10)
	root.Print()

	var pRoot *tree.Node
	pRoot.SetValue(110) //空结构体是nil

	pRoot = &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	root.Traverse()
}
