package main

import (
	"DSA/DS"
	"bufio"
	"fmt"
	"os"
)

var (
	reader *bufio.Reader = bufio.NewReader(os.Stdin)
	writer *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func main() {
	fmt.Println("Data structure loading...")
	tree := DS.InitTree(1)
	tree.LeftNode = &DS.Node{
		Data: 2,
		LeftNode: &DS.Node{
			Data: 4,
			LeftNode: &DS.Node{
				Data: 8,
				RightNode: &DS.Node{
					Data: 113,
				},
			},
			RightNode: &DS.Node{
				Data: 9,
				LeftNode: &DS.Node{
					Data: 111,
				},
				RightNode: &DS.Node{
					Data: 112,
				},
			},
		},
		RightNode: &DS.Node{
			Data: 5,
			LeftNode: &DS.Node{
				Data: 10,
			},
		},
	}
	tree.RightNode = &DS.Node{
		Data: 3,
		LeftNode: &DS.Node{
			Data: 6,
		},
		RightNode: &DS.Node{
			Data: 7,
			LeftNode: &DS.Node{
				Data: 19,
				LeftNode: &DS.Node{
					Data: 20,
				},
				RightNode: &DS.Node{
					Data: 100,
				},
			},
		},
	}
	//tree.PostOrderIteration(tree)
	//fmt.Println("Recursion")
	//tree.PostOrder(tree)

	//	fmt.Println(tree.MaxElement(tree))
	//	fmt.Println(tree.SearchElement(10))
	//	tree.LevelOrder()
	/*t := DS.InitTree(1)
	fmt.Println(t.InsertNode(2))
	fmt.Println(t.InsertNode(3))
	fmt.Println(t.InsertNode(4))
	fmt.Println(t.InsertNode(5))
	fmt.Println(t.InsertNode(6))
	fmt.Println(t.InsertNode(7))
	fmt.Println(t.InsertNode(8))
	fmt.Println(t.InsertNode(9))
	fmt.Println(t.InsertNode(10))
	*/
	//fmt.Println("No. of leaf node:", tree.LevelOrder())
	//fmt.Println("Tree Depth1:", t.MaxDeepNode())
	//fmt.Println("Tree Depth2:", tree.MaxDeepNode())
	//fmt.Println(tree.MaxLevelSum())
	//tree.PrintPath()
	//fmt.Println(tree.IsPathWithGivenSum(3))
	//fmt.Println(t.Size())
	//	inOrder := []string{"4", "8", "2", "5", "1", "6", "3", "7"}
	//	postOrder := []string{"8", "4", "5", "2", "6", "7", "3", "1"}
	//
	//	t2 := DS.PostInTree(postOrder, inOrder)
	//	t2.LevelOrder()
	//td := tree.FindLCAofNodes(8, 10)
	//if tree != nil {
	//	fmt.Println(td.Data)
	//} else {
	//	fmt.Println(td.Data)
	//}
	/*tree.VerticalSum()
	preorder := []string{"I", "L", "I", "L", "L"}
	t1 := DS.BuildTreeFromPreorder(preorder, 0)
	t1.LevelOrder()*/
	//parents := []string{"A", "B", "C", "*", "+", "D", "/"}
	// //fmt.Println(DS.FindHeightWithArray(parents))
	// tt := DS.BuildExpressionTree(parents)
	// tt.LevelOrder()
	bst := DS.InitTree(49)
	bst.LeftNode = &DS.Node{
		Data: 37,
		LeftNode: &DS.Node{
			Data: 13,
			LeftNode: &DS.Node{
				Data: 7,
			},
			RightNode: &DS.Node{
				Data: 19,
				RightNode: &DS.Node{
					Data: 25,
				},
			},
		},
		RightNode: &DS.Node{
			Data: 41,
		},
	}
	bst.RightNode = &DS.Node{
		Data: 89,
		LeftNode: &DS.Node{
			Data: 53,
			RightNode: &DS.Node{
				Data: 71,
				LeftNode: &DS.Node{
					Data: 60,
				},
				RightNode: &DS.Node{
					Data: 82,
				},
			},
		},
	}
	bst1 := DS.InitTree(3)
	bst1.LeftNode = &DS.Node{
		Data: 2,
	}
	bst1.RightNode = &DS.Node{
		Data: 7,
		LeftNode: &DS.Node{
			Data: 5,
			RightNode: &DS.Node{
				Data: 4,
			},
		},
	}
	// fmt.Println(bst1.MaxLevelSum())
	// q := DS.InitQueue(5)
	// node := bst.FindKthSmallElement(3, q)
	//fmt.Println(node.Data)
	// // // fmt.Println(bst.IsBST())
	// dummyNode := DS.InitLinkedList()
	// bst.BstToDLL(bst, bst, dummyNode)
	// dummyNode.Traverse()
	// node := DS.InitLinkedList()
	// node.Data = int32(1)
	// node.Next = &DS.ListNode{
	// 	Data: int32(10),
	// 	Next: &DS.ListNode{
	// 		Data: int32(13),
	// 		Next: &DS.ListNode{
	// 			Data: int32(15),
	// 			Next: &DS.ListNode{
	// 				Data: int32(18),
	// 				Next: &DS.ListNode{
	// 					Data: int32(20),
	// 					Next: &DS.ListNode{
	// 						Data: int32(22),
	// 					},
	// 				},
	// 			},
	// 		},
	// },
	// }
	// t1 := node.ListToBst()
	// t1.LevelOrder()
	//fmt.Println("Floor  ", bst.FindFloor(10).Data)
	//bst.PrintElementsInRangeBst(3, 9)
	//fmt.Println(bst.SameElementsBst(bst1))
	//	treeNode.InOrder()
	//	DS.ConstructAVLWihtRange(1, 1<<3).InOrder()
	// gt := DS.GenerateMinAVL(3)
	// //fmt.Println(gt.IsAVL())
	// gt.LevelOrder()
	// t := gt.RemoveNodeWithSingleChild()
	// fmt.Println("After removing nodes")
	// t.InOrder()
	// bst.InOrder()
	// bb := bst.PruneBst(17, 41)
	// fmt.Println("After removing nodes")
	// bb.InOrder()
	//val := [][]int64{{2, 5, 10, 11}, {3, 4, 15, 112}, {3, 4, 12, 23}, {1, 23, 133, 211}}
	// // t111 := DS.SegmentTree(val)
	// fmt.Println(t111.SegmetnTreeSumInRange(0, 7))
	// t111.SegmentTreeUpdateInRange(2, 2)
	// fmt.Println(t111.SegmetnTreeSumInRange(0, 7))
	//ht := DS.InitHeap(17)
	// ht.Insert(5)
	// ht.Insert(10)
	// ht.Insert(20)
	// ht.Insert(1)
	// ht.Insert(50)
	// ht.Insert(3)
	// ht.Delete()
	// bst.PostOrderIterationV2(bst)
	// ht.BuildHeap(val, 7)
	// fmt.Println(DS.HeapSort(val, 7))
	// fmt.Println(bst1.MaxWidth())
	// fmt.Println(ht.DeleteNode(3))
	// fmt.Println(ht.Arr)
	// fmt.Println(DS.KthSmallestElementV2(val, 6))
	// qu := DS.InitHeapQueue(7)
	// qu.Enqueue(7)
	// qu.Enqueue(2)
	// qu.Enqueue(34)
	// fmt.Println(qu.Dequeue())
	// fmt.Println(qu.Dequeue())
	// fmt.Println(qu.Dequeue())
	// fmt.Println(qu.Dequeue())
	//fmt.Println(ht.MergeKList(val, 4, 4, 16))
	// ds := DS.MakeSetWithFastFind(10)
	// fmt.Println(ds.Find(3))
	// ds.UnionByRank(1, 2)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(2, 3)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(3, 4)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(5, 6)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(7, 8)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(8, 9)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(5, 7)
	// fmt.Println(ds.Set)
	// ds.UnionByRank(1, 5)
	// fmt.Println(ds.Set)
	graph := DS.InitGraph(4, false)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.DFS()
	fmt.Println("------------------------>")
	graph.BFS()
}
