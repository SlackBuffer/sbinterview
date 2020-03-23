package linkedlist

type node struct {
	data *nodeData
	next *node
}

type nodeData struct {
	n int
}

func CreateLinkedList(data []int) *node {
	if data == nil {
		return nil
	}

	var head node
	pre := &head
	for _, d := range data {
		// 创建新节点；新节点的 next 的初始值为零值 nil
		currNode := &node{
			data: &nodeData{n: d},
		}
		// 将前一个节点与新节点链接起来
		pre.next = currNode
		// 将前一个节点更新为新建的节点
		pre = currNode
	}
	return &head
}

func CountNumWithinLinkedList(arr []int, target int) int {
	head := CreateLinkedList(arr)
	if head == nil {
		return 0
	}

	var count int
	for nd := head.next; ; {
		if nd.data.n == target {
			count++
		}
		if nd.next == nil {
			break
		}
		nd = nd.next
	}
	return count
}

// [ ] 删除节点，插入节点
