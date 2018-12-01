package DS

import "fmt"

type Graph struct {
	AdjMatrix   [][]bool
	VertexCount int64
	IsDirected  bool
	IsVisited   []bool
	InDegree    []int64
}

type ListGraph struct {
	Arr         []*ListNode
	VertexCount int64
}

func (graph *Graph) AddEdge(row, col int64) {
	if graph == nil {
		return
	}
	if row >= 0 && row < graph.VertexCount && col >= 0 && col < graph.VertexCount {
		graph.AdjMatrix[row][col] = true
		if !graph.IsDirected {
			graph.AdjMatrix[col][row] = true
		}
	}
}

func (graph *Graph) RemoveEdge(row, col int64) {
	if graph == nil {
		return
	}
	if row >= 0 && row < graph.VertexCount && col >= 0 && col < graph.VertexCount {
		graph.AdjMatrix[row][col] = false
		if !graph.IsDirected {
			graph.AdjMatrix[col][row] = false
		}
	}
}

func (graph *Graph) DFS() {
	if graph == nil {
		return
	}

	stk := InitStack(int32(graph.VertexCount * graph.VertexCount))
	graph.IsVisited[0] = true
	stk.Push(int32(0))
	fmt.Println(0)
	for !stk.IsEmpty() {
		val := stk.Read().(int32)
		vtx := graph.unvisitedNode(int64(val))
		if vtx == -1 {
			stk.Pop()
		} else {
			graph.IsVisited[vtx] = true
			stk.Push(int32(vtx))
			fmt.Println(vtx)
		}
	}
	graph.resetVisited()
}

func (graph *Graph) BFS() {
	if graph == nil {
		return
	}

	q := InitQueue(int32(graph.VertexCount * graph.VertexCount))
	graph.IsVisited[0] = true
	q.EnQueue(int32(0))
	for !q.IsEmpty() {
		val := q.DeQueue().(int32)
		graph.addUnvisitedNode(int64(val), q)
		fmt.Println(val)
	}
	graph.resetVisited()
}

func (graph *Graph) addUnvisitedNode(val int64, q *Queue) {
	var idx int64
	for idx = 0; idx < graph.VertexCount; idx++ {
		if graph.AdjMatrix[val][idx] && graph.IsVisited[idx] == false {
			q.EnQueue(int32(idx))
			graph.IsVisited[idx] = true
		}
	}
}

func (graph *Graph) unvisitedNode(val int64) int64 {
	var idx int64
	for idx = 0; idx < graph.VertexCount; idx++ {
		if graph.AdjMatrix[val][idx] && graph.IsVisited[idx] == false {
			return idx
		}
	}
	return -1
}

func (graph *Graph) resetVisited() {
	var idx int64
	for idx = 0; idx < graph.VertexCount; idx++ {
		graph.IsVisited[idx] = false
	}
}

func (graph *Graph) IsEdge(row, col int64) bool {
	if graph == nil {
		return false
	}
	if row >= 0 && row < graph.VertexCount && col >= 0 && col < graph.VertexCount {
		if !graph.IsDirected {
			return graph.AdjMatrix[col][row] && graph.AdjMatrix[row][col]
		}
		return graph.AdjMatrix[row][col]
	}
	return false
}

func (graph *Graph) TopologicalSort() (result []int64) {
	stk := InitStack(int32(graph.VertexCount))

	var idx int64
	for idx = 0; idx < graph.VertexCount; idx++ {
		if !graph.IsVisited[idx] {
			TopologicalSortUtil(graph, idx, stk)
		}
	}

	for !stk.IsEmpty() {
		val := stk.Pop().(int64)
		result = append(result, val)
	}

	return result
}

func TopologicalSortUtil(graph *Graph, vertex int64, stk *Stack) {
	if graph == nil {
		return
	}

	graph.IsVisited[vertex] = true

	unvisited := graph.unvisitedNode(vertex)
	for unvisited >= 0 {
		if !graph.IsVisited[vertex] {
			TopologicalSortUtil(graph, unvisited, stk)
		}
		unvisited = graph.unvisitedNode(vertex)
	}

	stk.Push(vertex)
}

func (graph *Graph) ShortestPathUnweighted(s int64) {

	if graph == nil {
		return
	}

	dist := make([]*int, graph.VertexCount)
	path := make([]*int64, graph.VertexCount)
	q := InitQueue(int32(graph.VertexCount * graph.VertexCount))

	for val := range dist {
		*dist[val] = -1
	}
	*dist[s] = 0

	q.EnQueue(int32(s))
	for !q.IsEmpty() {
		val := q.DeQueue().(int32)
		graph.addUnvisitedNodeToPath(int64(val), q, dist, path)
	}
	graph.resetVisited()

}

func (graph *Graph) addUnvisitedNodeToPath(val int64, q *Queue, dist []*int, path []*int64) {
	var idx int64
	for idx = 0; idx < graph.VertexCount; idx++ {
		if graph.AdjMatrix[val][idx] {
			if *dist[idx] == -1 {
				*dist[idx] = *dist[val] + 1
				*path[idx] = val
				q.EnQueue(int32(idx))
			}
		}
	}
}

func (graph *Graph) AddIndegree(indegree []int64) {
	if graph == nil {
		return
	}
	graph.InDegree = indegree
}

// Adjacency  Matrix imp.
func InitGraph(vertex int64, directed bool) *Graph {
	rows := make([][]bool, vertex)
	visited := make([]bool, vertex)
	indegree := make([]int64, vertex)
	for rw := range rows {
		cols := make([]bool, vertex)
		rows[rw] = cols
	}
	return &Graph{
		AdjMatrix:   rows,
		VertexCount: vertex,
		IsDirected:  directed,
		IsVisited:   visited,
		InDegree:    indegree,
	}
}

func InitLinkedListGraph(vertex int64, data []int64) *ListGraph {
	arr := make([]*ListNode, vertex)

	for idx := range data {
		arr[idx] = &ListNode{
			Data: data[idx],
		}
	}

	return &ListGraph{
		Arr:         arr,
		VertexCount: vertex,
	}
}
func (graph *ListGraph) AddEdges(source, destination int64) {
	if graph == nil {
		return
	}

	if source != -1 && destination != -1 {
		graph.Arr[source].pushBack(destination)
		graph.Arr[destination].pushBack(source)
	}
}

func (list *ListNode) pushBack(source int64) {
	if list.Next != nil {
		temp := list.Next
		list.Next = &ListNode{
			Data: source,
			Next: temp,
		}
	} else {
		list.Next = &ListNode{
			Data: source,
		}
	}
}

func getInt(data interface{}) int64 {
	if val, ok := data.(int64); ok {
		return val
	}
	return 0
}
