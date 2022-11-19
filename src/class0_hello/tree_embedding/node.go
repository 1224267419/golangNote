package tree_embedding

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node //左右树枝
}

func CreateNode(value int) *Node { //工厂函数,专门创建treeNode类型的变量
	return &Node{Value: value} //虽然是局部变量,仍然能被传出参数并使用,使用指针使得避免被垃圾回收
}

func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

func Print2(node Node) {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("这是空指针,不能传值")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() { //前序遍历
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
func (node *Node) postOrder() { //后序遍历
	if node == nil { //防止包装了空nide导致无法调出node的左子树
		return
	}

	node.Left.postOrder()
	node.Right.postOrder()
	node.Print()

}
