package DS

import (
	"fmt"
)

//------------------------------->>>>>Threaded Binary Tree<<<<----------------------------
//Left node will contain predecessor.
//Right node will contain successor.
//------------------------------------------END-------------------------------------------

type ThreadedNode struct {
	Data  int32
	Left  *ThreadedNode
	Right *ThreadedNode
	LTag  int32
	RTag  int32
}

type ThreadedTree interface {
}

func (tree *ThreadedNode) InorderSuccessor() *ThreadedNode {
	if tree == nil {
		return nil
	}
	if tree.RTag == 0 {
		return tree.Right
	} else {
		tempNode := tree.Right
		for tempNode.LTag == 1 {
			tempNode = tempNode.Left
		}
		return tempNode
	}
}

func (tree *ThreadedNode) PreorderSuccessor() *ThreadedNode {
	if tree == nil {
		return nil
	}
	if tree.LTag == 1 {
		return tree.Left
	} else {
		tempNode := tree
		for tempNode.RTag == 0 {
			tempNode = tempNode.Right
		}
		return tempNode.Right
	}
}

func (tree *ThreadedNode) InOrderTraversal() {
	tempNode := tree.InorderSuccessor()
	for tempNode != tree {
		tempNode = tree.InorderSuccessor()
		fmt.Println(tempNode.Data)
	}
}

func (tree *ThreadedNode) PreOrderTraversal() {
	tempNode := tree.PreorderSuccessor()
	for tempNode != tree {
		tempNode = tree.PreorderSuccessor()
		fmt.Println(tempNode.Data)
	}
}

func InitThreadedTree(data int32) *ThreadedNode {
	root := ThreadedNode{
		Data: data,
	}
	return &root
}
