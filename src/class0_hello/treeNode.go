package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode //左右树枝
}

func createNode(value int) *treeNode { //工厂函数,专门创建treeNode类型的变量
	return &treeNode{value: value} //虽然是局部变量,仍然能被传出参数并使用,使用指针使得避免被垃圾回收
}

func (node treeNode) print() {

	fmt.Println(node.value)
}

func print2(node treeNode) {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("这是空指针,不能传值")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() { //前序遍历
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.print()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) //new()新建对象
	//无论是地址还是结构本身,都用.来访问成员

	root.left.right = createNode(1)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.print()
	root.setValue(10)
	root.print()

	var pRoot *treeNode
	pRoot.setValue(110) //空结构体是nil

	pRoot = &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	root.traverse()
}
