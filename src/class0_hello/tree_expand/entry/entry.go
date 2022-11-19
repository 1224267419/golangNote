package main

import (
	"fmt"
	"tree_expand"
)

type myTreeNode struct {
	node *tree_expand.Node //指针引用创建,
}

func (myNode *myTreeNode) postOrder() { //后序遍历
	if myNode == nil || myNode.node == nil { //防止包装了空nide导致无法调出node的左子树
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree_expand.Node
	root = tree_expand.Node{Value: 3}
	root.Left = &tree_expand.Node{}
	root.Right = &tree_expand.Node{5, nil, nil}
	root.Right.Left = new(tree_expand.Node)
	root.Left.Right = tree_expand.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse() //前序遍历
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder() //后序遍历
}
