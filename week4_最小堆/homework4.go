package main

import (
	"fmt"
)

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i := len(nodes)/2 - 1; i >= 0; i-- {
		down(nodes, i, len(nodes))
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	current := i           // 需要下沉的元素, 值大于它的左右孩子
	child := 2*current + 1 // 左孩子
	temp := nodes[current].Value
	for child < n {
		if child+1 < n && nodes[child+1].Value < nodes[child].Value {
			child++
		}
		// 值小于两个孩子 则到达正确位置
		if temp <= nodes[child].Value {
			break
		}
		nodes[current].Value = nodes[child].Value
		current = child
		child = child*2 + 1
	}
	nodes[current].Value = temp
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	current := j
	parent := (j - 1) / 2
	for current > 0 && nodes[current].Value <= nodes[parent].Value {
		temp := nodes[current].Value
		nodes[current].Value = nodes[parent].Value
		nodes[parent].Value = temp
		current = parent
		parent = (parent - 1) / 2
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	min := nodes[0]
	nodes[0].Value = nodes[len(nodes)-1].Value
	nodes = nodes[:len(nodes)-1]
	down(nodes, 0, len(nodes)-1)
	return min, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i := 0; i < len(nodes); i++ {
		if node.Value == nodes[i].Value {
			nodes[i].Value = nodes[len(nodes)-1].Value
			nodes = nodes[0 : len(nodes)-1]
			down(nodes, 0, len(nodes)-1)
			break
		}
	}
	return nodes
}

func printHeap(nodes []Node) {
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("%d ", nodes[i].Value)
	}
	fmt.Printf("\n\n")
}

func main() {
	// 一组测试数据, Heap长度为8
	nodes := []Node{
		Node{9},
		Node{3},
		Node{10},
		Node{2},
		Node{4},
		Node{1},
		Node{5},
		Node{7},
		Node{6},
	}

	printHeap(nodes)

	fmt.Printf("Call function: Init()\n")
	Init(nodes)
	printHeap(nodes)

	fmt.Printf("Call function: up(), with adding 0\n")
	temp := Node{0}
	nodes = append(nodes, temp)
	up(nodes, 9)
	printHeap(nodes)

	fmt.Printf("Call function: Pop()\n")
	min, nodes := Pop(nodes)
	fmt.Printf("The minimum :%d\n\n", min)
	printHeap(nodes)

	fmt.Printf("Call function: Remove(), with removing 7\n")
	temp = Node{7}
	nodes = Remove(nodes, temp)
	printHeap(nodes)

	fmt.Printf("Call function: Push(), with pushing 8\n")
	temp = Node{8}
	nodes = Push(temp, nodes)
	printHeap(nodes)
}
