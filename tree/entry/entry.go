package main

import (
	"fmt"
	"learnGo/tree"
)

//组合来扩展
type MyTreeNode struct {
	node *tree.Node
}

//后序遍历
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()

	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
