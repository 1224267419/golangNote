package main

import (
	"fmt"
	"tree_embedding"
)

type myTreeNode struct {
	*tree_embedding.Node //embedding内嵌,语法糖,省下代码量
}

func (myNode *myTreeNode) postOrder() { //后序遍历
	if myNode == nil || myNode.Node == nil { //myNode.Node就是.出来的,用内嵌省下了结构体的别名
		return
	}

	left := myTreeNode{myNode.Left} //不再需要别名node
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()
	myNode.Print()

}
func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed")
	//不是重载而是shadowed,不过效果和重载很像
}

func main() {

	root := myTreeNode{&tree_embedding.Node{Value: 3}}
	//和tree_expand对比一下,直接用myTreeNode{}进行定义,而且后面仍能使用tree_embedding.Node的方法

	root.Left = &tree_embedding.Node{}
	root.Right = &tree_embedding.Node{5, nil, nil}
	root.Right.Left = new(tree_embedding.Node)
	root.Left.Right = tree_embedding.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse()      //用myTreeNode的方法
	root.Node.Traverse() //用回Node的方法

	fmt.Println()
	root.postOrder() //后序遍历
}
