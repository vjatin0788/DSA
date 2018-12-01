package DS

import (
	"fmt"
	"math"
	"regexp"
)

type Node struct {
	LeftNode   *Node
	RightNode  *Node
	Data       int32
	DataString string
	DataNode   interface{}
	Height     float64
	SizeOfTree int32
	LeftRange  int32
	RightRange int32
}

var (
	elementCount int32
)

type TreeImp interface {
	PreOrder()
	PreOrderIteration(root *Node)
	PostOrder()
	PostOrderIteration(root *Node)
	InOrder()
	InOrderIteration(root *Node)
	LevelOrder() int32
	MaxElement() int32
	SearchElement(data int32) bool
	MaxTreeDepth() int32
	MinTreeDepth() int32
	Size() int32
	Equals(tree2 *Node) bool
}

func (node *Node) MaxElement() int32 {
	//return MaxElementNode(root)
	return MaxElementNode(node)
}

func (node *Node) PreOrder() {
	preOrderTraversal(node)
}

func (node *Node) PostOrder() {
	postOrderTraversal(node)
}

func (node *Node) InOrder() {
	inOrderTraversal(node)
}

func (node *Node) Size() int32 {
	return sizeOf(node)
}

func (node *Node) MaxTreeDepth() int32 {
	return maxDepth(node)
}

func (node *Node) MinTreeDepth() int32 {
	return minDepth(node)
}
func (node *Node) FindAncestor(leaf *Node) {
	findAncestor(node, leaf)
}

func (node *Node) IsPathWithGivenSum(sum int32) bool {
	return findPath(node, 0, sum)
}

func PreInTree(preOrder, inOrder []string) *Node {
	if len(preOrder) == 0 || len(inOrder) != len(preOrder) {
		return nil
	}
	return preorderInorderTreeConstruction(0, len(inOrder)-1, 0, len(preOrder)-1, preOrder, inOrder)
}

func PostInTree(postOrder, inOrder []string) *Node {
	if len(postOrder) == 0 || len(inOrder) != len(postOrder) {
		return nil
	}
	return postorderInorderTreeConstruction(0, len(inOrder)-1, 0, len(postOrder)-1, postOrder, inOrder)
}
func (node *Node) MaxLevelSum() int32 {
	var max, i int32
	if node == nil {
		return 0
	}
	for i = 0; i < maxDepth(node); i++ {
		temp := levelSum(node, i)
		if temp > max {
			max = temp
		}
	}
	return max
}

func preOrderTraversal(root *Node) {
	if root != nil {
		fmt.Println(root.Data)
		preOrderTraversal(root.LeftNode)
		preOrderTraversal(root.RightNode)
	}
}

func postOrderTraversal(root *Node) {
	if root != nil {
		postOrderTraversal(root.LeftNode)
		postOrderTraversal(root.RightNode)
		fmt.Println(root.Data)
	}
}

func inOrderTraversal(root *Node) {
	if root != nil {
		inOrderTraversal(root.LeftNode)
		fmt.Println(root.Data)
		inOrderTraversal(root.RightNode)
	}
}

func (node *Node) PreOrderIteration(root *Node) {
	// add the data to tree
	//var temp interface{}
	stack := InitStack(5)

	if root == nil {
		fmt.Println("Empty Tree")
	}
	stack.Push(*root)

	for !stack.IsEmpty() {
		if nd, ok := stack.Pop().(Node); ok {
			fmt.Println(nd.Data)
			if nd.RightNode != nil {
				stack.Push(*nd.RightNode)
			}
			if nd.LeftNode != nil {
				stack.Push(*nd.LeftNode)
			}
		}
	}
}

func (node *Node) PostOrderIteration(root *Node) {
	// add the data to tree
	//var temp interface{}
	stack := InitStack(5)
	var prevNode *Node = nil

	if root == nil {
		fmt.Println("Empty Tree")
		return
	}
	stack.Push(root)

	for !stack.IsEmpty() {
		if nd, ok := stack.Read().(*Node); ok {
			if prevNode == nd.RightNode {
				fmt.Println(nd.Data)
				stack.Pop()
				prevNode = nd
				continue
			}
			if nd.LeftNode != nil && nd.LeftNode != prevNode {
				stack.Push(nd.LeftNode)
			} else if nd.RightNode != nil && nd.RightNode != prevNode {
				stack.Push(nd.RightNode)
			} else {
				fmt.Println(nd.Data)
				stack.Pop()
			}
			prevNode = nd
		}
	}
}

func (node *Node) PostOrderIterationV2(root *Node) {
	// add the data to tree
	//var temp interface{}
	stack := InitStack(20)
	var prevNode *Node = nil

	if root == nil {
		fmt.Println("Empty Tree")
		return
	}
	stack.Push(root)

	for !stack.IsEmpty() {
		if nd, ok := stack.Read().(*Node); ok {
			if nd == nil {
				stack.Pop()
			} else if nd != nil && prevNode == nd.RightNode {
				fmt.Println(nd.Data)
				stack.Pop()
			} else if nd != nil && nd.LeftNode != prevNode {
				stack.Push(nd.LeftNode)
			} else if nd != nil && nd.RightNode != prevNode {
				stack.Push(nd.RightNode)
			}
			prevNode = nd
		}
	}
}

func (node *Node) InOrderIteration(root *Node) {
	// add the data to tree
	//var temp interface{}
	// stack will be ade dynamic acc to size of tree in future
	stack := InitStack(5)
	var prevNode *Node = nil

	if root == nil {
		fmt.Println("Empty Tree")
	}
	stack.Push(root)

	for !stack.IsEmpty() {
		if nd, ok := stack.Read().(*Node); ok {
			if prevNode == nd.RightNode {
				stack.Pop()
				prevNode = nd
				continue
			} else if prevNode == nd.LeftNode {
				fmt.Println(nd.Data)
			}
			if nd.LeftNode != nil && nd.LeftNode != prevNode {
				stack.Push(nd.LeftNode)
			} else if nd.RightNode != nil && nd.RightNode != prevNode {
				stack.Push(nd.RightNode)
			} else {
				fmt.Println(nd.Data)
				stack.Pop()
			}
			prevNode = nd
		}
	}
}
func (node *Node) LevelOrder() int32 {
	// size will be made dynamic acc to size of tree
	q := InitQueue(20)
	var count int32
	if node != nil {
		q.EnQueue(node)
		for !q.IsEmpty() {
			if nd, ok := q.DeQueue().(*Node); ok {
				fmt.Println(nd.Data, nd.DataString)
				if nd.LeftNode != nil {
					q.EnQueue(nd.LeftNode)
				}
				if nd.RightNode != nil {
					q.EnQueue(nd.RightNode)
				}
				if nd.RightNode == nil && nd.LeftNode == nil {
					count++
				}
			}
		}
	}
	return count
}

func MaxElementNode(root *Node) int32 {
	if root == nil {
		return 0
	}
	return maxOf(root.Data, MaxElementNode(root.LeftNode), MaxElementNode(root.RightNode))
}

func maxOf(a, b, c int32) int32 {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func (node *Node) SearchElement(data int32) bool {
	q := InitQueue(10)
	if node != nil {
		q.EnQueue(node)
		for !q.IsEmpty() {
			if nd, ok := q.DeQueue().(*Node); ok {
				if data == nd.Data {
					return true
				}
				if nd.LeftNode != nil {
					q.EnQueue(nd.LeftNode)
				}
				if nd.RightNode != nil {
					q.EnQueue(nd.RightNode)
				}
			}
		}
	}
	return false
}

func (node *Node) InsertNode(data int32) bool {
	q := InitQueue(10)
	if node != nil {
		q.EnQueue(node)
		for !q.IsEmpty() {
			if nd, ok := q.DeQueue().(*Node); ok {
				if nd.LeftNode == nil {
					nd.LeftNode = &Node{
						Data: data,
					}
					return true
				} else if nd.RightNode == nil {
					nd.RightNode = &Node{
						Data: data,
					}
					return true
				}
				if nd.LeftNode != nil {
					q.EnQueue(nd.LeftNode)
				}
				if nd.RightNode != nil {
					q.EnQueue(nd.RightNode)
				}
			}
		}
	}
	return false
}

func sizeOf(root *Node) int32 {
	if root == nil {
		return 0
	}
	return 1 + sizeOf(root.LeftNode) + sizeOf(root.RightNode)
}

func maxDepth(root *Node) int32 {
	if root == nil {
		return 0
	}
	leftSub := maxDepth(root.LeftNode)
	rightSub := maxDepth(root.RightNode)

	if leftSub >= rightSub {
		return leftSub + 1
	} else {
		return rightSub + 1
	}
}

func minDepth(root *Node) int32 {
	if root == nil {
		return 0
	}
	leftSub := minDepth(root.LeftNode)
	rightSub := minDepth(root.RightNode)

	if leftSub >= rightSub {
		return rightSub + 1
	} else {
		return leftSub + 1
	}
}

func (node *Node) MaxDeepNode() int32 {
	q := InitQueue(10)
	if node != nil {
		q.EnQueue(node)
		for !q.IsEmpty() {
			if nd, ok := q.DeQueue().(*Node); ok {
				if nd.LeftNode != nil {
					q.EnQueue(nd.LeftNode)
				}
				if nd.RightNode != nil {
					q.EnQueue(nd.RightNode)
				}
				if q.IsEmpty() {
					return nd.Data
				}
			}
		}
	}
	return 0
}

func (node *Node) Equals(tree *Node) bool {
	if node == nil && tree == nil {
		return true
	}
	if node != nil && tree != nil {
		return node.Data == tree.Data && node.LeftNode.Equals(tree.LeftNode) && node.RightNode.Equals(tree.RightNode)
	}
	return false
}

func LeafNodeV2(root *Node, depth int32) int32 {
	if root == nil {
		return 0
	} else {
		if depth == 0 {
			return 1
		} else {
			return LeafNodeV2(root.LeftNode, depth-1) + LeafNodeV2(root.RightNode, depth-1)
		}
	}
}

//using preorder traversal
var Res []int32

func (node *Node) LevelWithMaxNodes() []int32 {
	Res = make([]int32, 20)
	binaryTreeLevelWithMaxNodes(node, 0)
	return Res
}
func binaryTreeLevelWithMaxNodes(root *Node, level int32) {
	if root == nil {
		return
	}

	Res[level] += 1
	binaryTreeLevelWithMaxNodes(root.LeftNode, level+1)
	binaryTreeLevelWithMaxNodes(root.RightNode, level+1)
}

var TreeDiameter float64

func (node *Node) TreeDiameter() float64 {
	node.diameter()
	return TreeDiameter
}

func (node *Node) diameter() float64 {
	if node == nil {
		return 0
	}
	left := node.LeftNode.diameter()
	right := node.RightNode.diameter()
	if left+right+1 > TreeDiameter {
		TreeDiameter = left + right + 1
	}
	return math.Max(left, right) + 1
}

func levelSum(root *Node, depth int32) int32 {
	if root == nil {
		return 0
	} else {
		if depth == 0 {
			return root.Data
		} else {
			return levelSum(root.LeftNode, depth-1) + levelSum(root.RightNode, depth-1)
		}
	}
}

func (node *Node) LevelSum(root *Node, size int32) int32 {
	if root == nil {
		return 0
	}
	q1 := InitQueue(size)
	q2 := InitQueue(size)
	var max, sum int32

	q1.EnQueue(root)
	q2.EnQueue(root)
	sum = root.Data
	for !q1.IsEmpty() {
		n1, _ := q1.DeQueue().(*Node)
		q2.DeQueue()
		if n1.LeftNode != nil {
			q1.EnQueue(n1.LeftNode)
		}
		if n1.RightNode != nil {
			q1.EnQueue(n1.RightNode)
		}
		sum += n1.Data
		if q2.IsEmpty() {
			q2.QueueBody = q1.QueueBody
			q2.Size = q1.Size
			q2.Front = q1.Front
			q2.Rear = q1.Rear
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}
	return max
}

func (node *Node) PrintPath() {
	arr := make([]int32, 100)
	printAllPath(node, arr, 0)
}

func printAllPath(root *Node, arr []int32, pathLen int) {

	if root == nil {
		return
	}
	arr[pathLen] = root.Data
	pathLen++
	if root.LeftNode == nil && root.RightNode == nil {
		for i := 0; i < pathLen; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println()
	} else {
		printAllPath(root.LeftNode, arr, pathLen)
		printAllPath(root.RightNode, arr, pathLen)
	}
}

//we can even minus the sum with data and check wheather final deducted sum is equal to leaf node or not!
func findPath(root *Node, pathSum, sum int32) bool {
	if root == nil {
		return false
	}
	sumVal := root.Data + pathSum

	if root.LeftNode == nil && root.RightNode == nil && sumVal == sum {
		return true
	} else {
		return findPath(root.LeftNode, sumVal, sum) || findPath(root.RightNode, sumVal, sum)
	}

}

func preorderInorderTreeConstruction(inStart, inEnd, preStart, preEnd int, preString, inString []string) *Node {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	var foundIndex int
	for i := inStart; i <= inEnd; i++ {
		if inString[i] == preString[preStart] {
			foundIndex = i
			break
		}
	}
	fmt.Printf("inStart: %d inEnd: %d preStart: %d preEnd: %d foundIndex: %d \n", inStart, inEnd, preStart, preEnd, foundIndex)
	left := preorderInorderTreeConstruction(inStart, foundIndex-1, preStart+1, preStart+foundIndex-inStart, preString, inString)
	rigth := preorderInorderTreeConstruction(foundIndex+1, inEnd, preStart+foundIndex-inStart+1, preEnd, preString, inString)

	node := Node{
		DataString: inString[foundIndex],
		LeftNode:   left,
		RightNode:  rigth,
	}
	return &node
}

//need more work on it... under construction.
func postorderInorderTreeConstruction(inStart, inEnd, postStart, postEnd int, postString, inString []string) *Node {
	if postStart > postEnd || inStart > inEnd {
		return nil
	}
	var foundIndex int
	for i := inStart; i <= inEnd; i++ {
		if inString[i] == postString[postEnd] {
			foundIndex = i
			break
		}
	}
	fmt.Printf("inStart: %d inEnd: %d preStart: %d preEnd: %d foundIndex: %d \n", inStart, inEnd, postStart, postEnd, foundIndex)
	left := preorderInorderTreeConstruction(inStart, foundIndex-1, postStart, postStart+foundIndex-inStart-1, postString, inString)
	rigth := preorderInorderTreeConstruction(foundIndex+1, inEnd, postStart+foundIndex-inStart, postEnd-1, postString, inString)

	node := Node{
		DataString: inString[foundIndex],
		LeftNode:   left,
		RightNode:  rigth,
	}
	return &node
}

func findAncestor(root, node *Node) bool {
	if root == nil {
		return false
	} else {
		if root.LeftNode == node || root.RightNode == node || findAncestor(root.LeftNode, node) || findAncestor(root.RightNode, node) {
			fmt.Println(root.Data)
			return true
		}
	}
	return false
}

func (node *Node) FindLCAofNodes(n1, n2 int32) int32 {
	b := findElement(node, n1)
	c := findElement(node, n2)

	if b && c {
		res := findLCA(node, n1, n2)
		if node != nil {
			return res.Data
		}
	}
	return -1
}

func findLCA(root *Node, node1, node2 int32) *Node {
	if root == nil {
		return nil
	}
	if root.Data == node1 || root.Data == node2 {
		return root
	}

	left := findLCA(root.LeftNode, node1, node2)
	right := findLCA(root.RightNode, node1, node2)

	if left != nil && right != nil {
		return root
	} else {
		if left != nil {
			return left
		} else {
			return right
		}
	}
}

func findElement(root *Node, a int32) bool {
	if root == nil {
		return false
	}

	if root.Data == a {
		return true
	}

	return findElement(root.LeftNode, a) || findElement(root.RightNode, a)
}

func (node *Node) ZigZag(size int32) {

	if node == nil {
		return
	}
	q := InitQueue(size)
	s1 := InitStack(size)
	s2 := InitStack(size)
	s3 := InitStack(size)
	q.EnQueue(node)
	for !q.IsEmpty() || !s2.IsEmpty() {
		if !q.IsEmpty() {
			val, ok := q.DeQueue().(*Node)
			if ok {
				fmt.Println(val.Data)
				if val.RightNode != nil {
					s1.Push(val.RightNode)
					if val.RightNode.LeftNode != nil {
						s2.Push(val.RightNode.RightNode)
					}
					if val.RightNode.RightNode != nil {
						s2.Push(val.RightNode.LeftNode)
					}
				}
				if val.LeftNode != nil {
					s1.Push(val.LeftNode)
					if val.LeftNode.RightNode != nil {
						s2.Push(val.LeftNode.RightNode)
					}
					if val.LeftNode.LeftNode != nil {
						s2.Push(val.LeftNode.LeftNode)
					}
				}
			}
		} else {
			for !s1.IsEmpty() {
				val := s1.Pop().(*Node)
				fmt.Println(val.Data)
			}
			for !s2.IsEmpty() {
				s3.Push(s2.Pop())
			}
			for !s3.IsEmpty() {
				q.EnQueue(s3.Pop())
			}
		}

	}
}

func (node *Node) VerticalSum() {
	if node == nil {
		return
	}
	hash := make(map[int32]int32)
	verticalSumBinaryTree(node, hash, 0)
	for k, v := range hash {
		fmt.Println(k, v)
	}
}

func verticalSumBinaryTree(root *Node, hash map[int32]int32, pos int32) {
	if root == nil {
		return
	}
	verticalSumBinaryTree(root.LeftNode, hash, pos-1)
	verticalSumBinaryTree(root.RightNode, hash, pos+1)

	if _, ok := hash[pos]; ok {
		hash[pos] += root.Data
	} else {
		hash[pos] = root.Data
	}
}

func BuildTreeFromPreorder(preorder []string, idx int) *Node {
	return buildFromPreorder(preorder, idx)
}

func buildFromPreorder(preorder []string, idx int) *Node {
	if len(preorder) == 0 || len(preorder) == idx {
		return nil
	}
	var node *Node
	node = &Node{
		DataString: preorder[idx],
	}
	//check for Leaf
	if preorder[idx] == "L" {
		return node
	}
	idx++
	node.LeftNode = buildFromPreorder(preorder, idx)
	idx++
	node.RightNode = buildFromPreorder(preorder, idx)
	return node
}

func FindHeightWithArray(parents []int32) int32 {

	var (
		idx      int32
		idy      int32
		maxDepth int32
		curDepth int32
	)
	for idx = 0; idx < int32(len(parents)); idx++ {
		curDepth = 0
		idy = idx
		for parents[idy] != -1 {
			curDepth++
			idy = parents[idy]
		}
		if curDepth > maxDepth {
			maxDepth = curDepth
		}
	}
	return maxDepth
}

func (root *Node) InorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node == rightMostNode(root) {
		return nil
	} else if node.RightNode != nil {
		return findInRightNode(node)
	} else {
		return findParentInorderSuccessor(root, node)
	}
}

func findInRightNode(root *Node) *Node {
	if root.LeftNode != nil {
		findInRightNode(root.LeftNode)
	}
	return root
}

func rightMostNode(root *Node) *Node {
	temp := root.RightNode
	for temp != nil {
		temp = temp.RightNode
	}
	return temp
}

func findParentInorderSuccessor(root, node *Node) *Node {
	if root == nil {
		return nil
	}
	if root == node {
		return root
	}

	left := findParentInorderSuccessor(root.LeftNode, node)
	right := findParentInorderSuccessor(root.RightNode, node)

	if root.LeftNode == right || root.LeftNode == left {
		fmt.Println("successor is:", root)
		return nil
	} else {
		return root
	}
}

func BuildExpressionTree(expression []string) *Node {

	if expression == nil {
		return nil
	}
	//init stack
	s1 := InitStack(int32(len(expression)))

	for _, str := range expression {
		//check whether element is operand than add it to the stack
		if match, _ := regexp.MatchString("[A-Z]", str); match {
			node := &Node{
				DataString: str,
			}
			s1.Push(node)
		}
		if match, _ := regexp.MatchString("[^A-Za-z0-9_]", str); match {
			node := &Node{
				DataString: str,
			}

			right := s1.Pop()
			if val, ok := right.(*Node); ok {
				node.RightNode = val
			}

			left := s1.Pop()
			if val, ok := left.(*Node); ok {
				node.LeftNode = val
			}

			s1.Push(node)
		}
	}

	node := s1.Pop()
	if val, ok := node.(*Node); ok {
		return val
	}

	return nil
}

func (node *Node) BinarySearch(data int32) *Node {
	return binarySearch(node, data)
}

func binarySearch(root *Node, data int32) *Node {
	if root == nil {
		return nil
	}
	if root.Data > data {
		return binarySearch(root.LeftNode, data)
	}
	if root.Data < data {
		return binarySearch(root.RightNode, data)
	}
	return root
}

func (node *Node) BinaryMin() *Node {
	if node == nil {
		return nil
	}
	if node.LeftNode != nil {
		return node.LeftNode.BinaryMin()
	}
	return node
}

func (node *Node) BinaryMax() *Node {
	if node == nil {
		return nil
	}
	if node.RightNode != nil {
		return node.RightNode.BinaryMin()
	}
	return node
}

func (node *Node) InsertBST(data int32) {
	if node == nil {
		return
	}
	if data < node.Data {
		if node.LeftNode != nil {
			node.LeftNode.InsertBST(data)
		} else {
			tempNode := &Node{Data: data}
			node.LeftNode = tempNode
			return
		}
	}
	if data > node.Data {
		if node.RightNode != nil {
			node.RightNode.InsertBST(data)
		} else {
			tempNode := &Node{Data: data}
			node.RightNode = tempNode
			return
		}
	}

}

func (node *Node) DeleteBST(data int32, prev *Node) {
	if node == nil {
		return
	}
	if node.Data == data {
		if node.LeftNode != nil && node.RightNode != nil {
			max := node.LeftNode.BinaryMax()
			node.LeftNode.DeleteBST(max.Data, node)
			prev = addNodeToPrevNode(prev, max, data)
		} else if node.LeftNode != nil {
			prev = addNodeToPrevNode(prev, node.LeftNode, data)
		} else if node.RightNode != nil {
			prev = addNodeToPrevNode(prev, node.RightNode, data)
		} else {
			prev = addNodeToPrevNode(prev, nil, data)
		}
	}
	if data < node.Data {
		node.LeftNode.DeleteBST(data, node)
	}
	if data > node.Data {
		node.RightNode.DeleteBST(data, node)
	}
}

func addNodeToPrevNode(prev, node *Node, data int32) *Node {
	if prev == nil {
		return node
	}
	if prev.LeftNode.Data == data {
		prev.LeftNode = node
	}
	if prev.RightNode.Data == data {
		prev.RightNode = node
	}
	return prev
}

func (node *Node) FindLcaBST(a, b float64) *Node {
	if node == nil {
		return nil
	}
	if int32(math.Min(a, b)) > node.Data {
		return node.RightNode.FindLcaBST(a, b)
	}
	if int32(math.Max(a, b)) < node.Data {
		return node.LeftNode.FindLcaBST(a, b)
	}
	return node
}

func (node *Node) IsBST() bool {
	if node == nil {
		return true
	}
	if node.LeftNode != nil && node.LeftNode.BinaryMax().Data > node.Data {
		return false
	}
	if node.RightNode != nil && node.RightNode.BinaryMax().Data < node.Data {
		return false
	}
	return node.LeftNode.IsBST() && node.RightNode.IsBST()
}

func (node *Node) IsBSTv2(min, max int32) bool {
	if node == nil {
		return true
	}
	return min < node.Data && node.Data < max && node.LeftNode.IsBSTv2(min, node.Data) && node.RightNode.IsBSTv2(node.Data, max)
}

//Third way of implementing IsBst is to do inorder traversal as it will produce sorted list and check current val should be greater than prev val.

func (node *Node) BstToDLL(root, prev *Node, head *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	//process the left node
	left := node.LeftNode.BstToDLL(root, node, head)
	listNode := &ListNode{
		Data: node.Data,
		Prev: left,
	}
	if left != nil {
		left.Next = listNode
	}
	right := node.RightNode.BstToDLL(root, node, head)
	listNode.Next = right

	if node == root.BinaryMin() {
		head.Next = listNode
	}
	if node == root.BinaryMax() {
		listNode.Next = head.Next
		head.Next.Prev = listNode
	}

	if prev.RightNode == node {
		if left == nil {
			return listNode
		}
		return left
	}
	if prev.LeftNode == node {
		if right == nil {
			return listNode
		}
		return right
	}
	return nil
}

func (node *Node) FindKthSmallElement(k int32, q *Queue) *Node {
	if node == nil {
		return nil
	}
	left := node.LeftNode.FindKthSmallElement(k, q)
	if left != nil {
		return left
	}
	q.EnQueue(node.Data)
	if q.Size == k {
		return node
	}
	return node.RightNode.FindKthSmallElement(k, q)
}

func (node *Node) FindCeil(key int32) *Node {
	if node == nil {
		return nil
	}
	return ceilBst(node, key)
}

func ceilBst(root *Node, key int32) *Node {
	if root == nil {
		return nil
	}

	left := ceilBst(root.LeftNode, key)
	if left != nil {
		return left
	}
	if root.Data >= key {
		return root
	}
	return ceilBst(root.RightNode, key)
}

func (node *Node) FindFloor(key int32) *Node {
	if node == nil {
		return nil
	}
	size := sizeOf(node)
	stk := InitStack(size)
	return floorBst(node, stk, key, size)
}

func floorBst(root *Node, stk *Stack, key, size int32) *Node {
	if root == nil {
		return nil
	}

	left := floorBst(root.LeftNode, stk, key, size)
	stk.Push(root)
	if left != nil {
		return left
	} else {
		fmt.Println(stk.size, root.Data, size)
		if stk.size == size {
			val := stk.Read().(*Node)
			return val
		}
	}
	if root.Data == key {
		return root
	}
	if root.Data > key {
		if stk.size == 1 {
			return &Node{}
		}
		stk.Pop()
		val := stk.Read().(*Node)
		return val
	}
	return floorBst(root.RightNode, stk, key, size)
}

/*finding union and intersection of bst
A) 1)convert into 2 sorted arrays.
   2)find the common elemet/merge the arrays.

B) 1)Convert into dll both the list
   2)find the common/merge the list
   3)convert into bst.
C) 1)traverse inorder bst and store keys in hash
   2)traverse 2nd bst and check whether the hash exist or not
*/

func (node *Node) PrintElementsInRangeBst(k1, k2 int32) {
	if node == nil {
		return
	}
	if node.Data >= k1 {
		node.LeftNode.PrintElementsInRangeBst(k1, k2)
	}
	if node.Data >= k1 && node.Data <= k2 {
		fmt.Println(node.Data)
	}
	if node.Data <= k2 {
		node.RightNode.PrintElementsInRangeBst(k1, k2)
	}
}

/* Alternate way
---->Threaded trees can be used to go to inorder successor
First find the element greater than k1 and than use inorder successor to check each element unless you encounter the k2.

---->Using Level Order Traversal.
*/

func (node *Node) SameElementsBst(n1 *Node) bool {
	if node == nil || n1 == nil {
		return false
	}
	val := make(map[int32]bool)
	fillMap(node, val)

	return compareMap(n1, val)
}

func fillMap(root *Node, val map[int32]bool) {
	if root == nil {
		return
	}
	fillMap(root.LeftNode, val)
	val[root.Data] = true
	fillMap(root.RightNode, val)
}

func compareMap(root *Node, val map[int32]bool) bool {
	if root == nil {
		return true
	}
	left := compareMap(root.LeftNode, val)
	right := compareMap(root.RightNode, val)

	_, ok := val[root.Data]
	if !ok {
		return false
	}
	if left && right {
		return true
	}
	return false
}

/*------------>AVL Trees<--------------
1) It is Bst
2) HB(1) , Height balance
3) Max Nodes N(H) = N(H-1) + N(H-2) + 1, wheras N(0)=1, N(1)=2
4) Min Nodes N(H) = 2*N(H-1) + 1,
5) Min Height with nodes n, log2 n.
6) Max Height with nodes n, 1.44*log2 n.
*/

func (node *Node) getHeight() float64 {
	if node == nil {
		return -1
	}
	return node.Height
}

func leftRotate(root *Node) *Node {
	if root == nil {
		return nil
	}
	temp := root.LeftNode
	root.LeftNode = temp.RightNode
	temp.RightNode = root

	//set height
	root.Height = math.Max(root.LeftNode.getHeight(), root.RightNode.getHeight()) + 1
	temp.Height = math.Max(root.Height, root.LeftNode.getHeight()) + 1
	return temp
}

func rightRotate(root *Node) *Node {
	if root == nil {
		return nil
	}
	temp := root.RightNode
	root.RightNode = temp.LeftNode
	temp.LeftNode = root

	root.Height = math.Max(root.LeftNode.getHeight(), root.RightNode.getHeight()) + 1
	temp.Height = math.Max(root.Height, root.RightNode.getHeight()) + 1

	return temp
}

func doubleRotateRightLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.LeftNode = rightRotate(root.LeftNode)
	return leftRotate(root)
}

func doubleRotateLeftRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.RightNode = leftRotate(root.RightNode)
	return rightRotate(root)
}

func (node *Node) InsertAVL(data int32) *Node {
	if node == nil {
		return nil
	}
	return insertAVL(node, data)
}

func insertAVL(root *Node, data int32) *Node {
	if root == nil {
		return &Node{
			Data:   data,
			Height: 0,
		}
	}
	if data < root.Data {
		root.LeftNode = insertAVL(root.LeftNode, data)
		if root.LeftNode.getHeight()-root.RightNode.getHeight() == 2 {
			if root.LeftNode.Data > data {
				return leftRotate(root)
			}
			return doubleRotateRightLeft(root)
		}
	}
	if data > root.Data {
		root.RightNode = insertAVL(root.RightNode, data)
		if root.RightNode.getHeight()-root.LeftNode.getHeight() == 2 {
			if root.RightNode.Data < data {
				return rightRotate(root)
			}
			return doubleRotateLeftRight(root)
		}
	}

	root.Height = math.Max(root.LeftNode.getHeight(), root.RightNode.getHeight()) + 1
	return root
}

func ConstructAVLWithZero(height int32) *Node {
	if height < 0 {
		return nil
	}
	node := &Node{}
	node.LeftNode = ConstructAVLWithZero(height - 1)
	node.Data = elementCount
	elementCount++
	node.RightNode = ConstructAVLWithZero(height - 1)

	return node
}

//initial call  can be ConstructAVLWihtRange(1,1<<h) ...as left shift will make the numer 2^h
func ConstructAVLWihtRange(left, right int32) *Node {
	if left > right {
		return nil
	}
	var mid int32

	mid = (right + left) / 2

	node := &Node{}
	node.LeftNode = ConstructAVLWihtRange(left, mid-1)
	node.Data = mid
	node.RightNode = ConstructAVLWihtRange(mid+1, right)

	return node
}

func (node *Node) IsAVL() bool {
	bf := isAVL(node)
	if math.Abs(bf) <= 1 {
		return true
	}
	return false
}
func isAVL(node *Node) float64 {
	if node == nil {
		return node.getHeight()
	}

	left := isAVL(node.LeftNode)
	right := isAVL(node.RightNode)

	return left - right
}

func GenerateMinAVL(height int32) *Node {
	if height < 0 {
		return nil
	}
	var root *Node
	root = &Node{}

	root.LeftNode = GenerateMinAVL(height - 1)
	root.Data = elementCount
	elementCount++
	root.RightNode = GenerateMinAVL(height - 2)
	root.Height = root.LeftNode.getHeight() + 1

	return root
}

func (node *Node) AvlRange(a, b int32) int32 {
	if node == nil {
		return 0
	}
	if node.Data > b {
		return node.LeftNode.AvlRange(a, b)
	}
	if node.Data < a {
		return node.RightNode.AvlRange(a, b)
	}
	return node.LeftNode.AvlRange(a, b) + node.RightNode.AvlRange(a, b) + 1
}

func (node *Node) RemoveNodeWithSingleChild() *Node {
	if node == nil {
		return nil
	}
	return removeNode(node)
}

func removeNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.LeftNode = removeNode(root.LeftNode)
	root.RightNode = removeNode(root.RightNode)

	if root.LeftNode != nil && root.RightNode == nil {
		return root.LeftNode
	}
	if root.RightNode != nil && root.LeftNode == nil {
		return root.RightNode
	}

	return root
}

func (node *Node) PruneBst(min, max int32) *Node {
	if node == nil {
		return nil
	}
	return pruneBst(node, min, max)
}

func pruneBst(root *Node, min, max int32) *Node {
	if root == nil {
		return nil
	}

	if root.Data > min {
		root.LeftNode = pruneBst(root.LeftNode, min, max)
	} else {
		root.LeftNode = nil
	}

	if root.Data < max {
		root.RightNode = pruneBst(root.RightNode, min, max)
	} else {
		root.RightNode = nil
	}

	if root.Data >= min && root.Data <= max {
		return root
	}

	if root.LeftNode != nil {
		return root.LeftNode
	} else {
		return root.RightNode
	}

	return nil
}

/*
-> we can traverse the level order without queue by linking the current level nodes and loop over it.
Algorithm for calculating leaf nodes
Algorithm Traverse(r){
	if r is leaf return 1
	else{
		a = Traverse(r.LeftNode)
		b = Traverse(r.RightNode)
		return a+b
	}
}
*/

func getTreeSize(root *Node) int32 {
	if root == nil {
		return 0
	}
	return root.SizeOfTree
}

func (node *Node) KthSmallest(k int32) *Node {
	if node == nil {
		return nil
	}
	r := getTreeSize(node.LeftNode) + 1

	if r == k {
		return node
	}
	if r > k {
		return node.LeftNode.KthSmallest(k)
	} else {
		return node.RightNode.KthSmallest(k - r)
	}
}

func SegmentTree(nums []int32) *Node {
	if nums == nil {
		return nil
	}
	var height, idx, actualHeight int32
	height = 0
	idx = 1

	for idx < int32(len(nums)) {
		idx <<= 1
		actualHeight++
	}
	return segmentTreeBuilder(height, actualHeight, nums)
}

// Segment tree can be constructed acc to requirement. ---------------------------------><<<<<<<<<<<
func segmentTreeBuilder(height, actualHeight int32, nums []int32) *Node {
	if height > actualHeight {

		var val int32
		val = 0
		if elementCount < int32(len(nums)) {
			val = nums[elementCount]
		}
		node := &Node{
			Data:       val,
			LeftRange:  elementCount,
			RightRange: elementCount,
		}

		elementCount++

		return node
	}

	left := segmentTreeBuilder(height+1, actualHeight, nums)
	right := segmentTreeBuilder(height+1, actualHeight, nums)

	node := &Node{
		Data:       left.Data + right.Data,
		LeftRange:  left.LeftRange,
		RightRange: right.RightRange,
		LeftNode:   left,
		RightNode:  right,
	}

	return node

}

func (node *Node) SegmetnTreeSumInRange(left, right int32) int32 {
	if node == nil {
		return 0
	}
	if node.RightRange < left || node.LeftRange > right {
		return 0
	}
	if node.LeftRange >= left && node.RightRange <= right {
		return node.Data
	}
	return node.LeftNode.SegmetnTreeSumInRange(left, right) + node.RightNode.SegmetnTreeSumInRange(left, right)

}

func (node *Node) SegmentTreeUpdateInRange(num, index int32) *Node {
	if node == nil {
		return nil
	}
	if node.LeftRange == num && node.RightRange == num {
		node.Data = num
		return node
	}
	if node.LeftRange == node.RightRange {
		return node
	}

	left := node.LeftNode.SegmentTreeUpdateInRange(num, index)
	right := node.RightNode.SegmentTreeUpdateInRange(num, index)

	node.Data = left.Data + right.Data
	return node
}

//------------------->>>>>>>>>>>>>>Segment  tree ends <<<<<<<<<<<<--------

func InitTree(data int32) *Node { //to implement interface return TreeImp
	root := Node{
		Data: data,
	}
	return &root
}

/* Points to remember for trees.
-> Check for single node and it's left and right subtree. same thing will happen across the tree.
-> Median is root for odd number of integer and root,inorder successor for root for even number if we use AVL.
-> A complete binary tree is more efficient than balanced binary tree  because balanced binary tree use more space to store some flages. Eg. Red/Black uses bit to store flag and
   binary tree can be save in array which will not use pointer and no ram is used. A pointer uses RAM.
*/

/* -----------------------RED BLACK----------------------
1) Every node has a color either red or black.

2) Root of tree is always black.

3) There are no two adjacent red nodes (A red node cannot have a red parent or red child).

4) Every path from root to a NULL node has same number of black nodes.

Black Height of a Red-Black Tree :
Black height is number of black nodes on a path from a node to a leaf. Leaf nodes are also counted black nodes. From above properties 3 and 4, we can derive, a node of height h has black-height >= h/2.


Every Red Black Tree with n nodes has height <= 2Log2(n+1)

This can be proved using following facts:
1) For a general Binary Tree, let k be the minimum number of nodes on all root to NULL paths, then n >= 2k – 1 (Ex. If k is 3, then n is atleast 7). This expression can also be written as k <= 2Log2(n+1)

2) From property 4 of Red-Black trees and above claim, we can say in a Red-Black Tree with n nodes, there is a root to leaf path with at-most Log2(n+1) black nodes.

3) From property 3 of Red-Black trees, we can claim that the number black nodes in a Red-Black tree is at least ⌊ n/2 ⌋ where n is the total number of nodes.

From above 2 points, we can conclude the fact that Red Black Tree with n nodes has height <= 2Log2(n+1)
*/
