package linked_list

import (
	"fmt"
	"sync"
)

type Node struct {
	data int64
	prev *Node
	next *Node
}

type List struct {
	mutex    *sync.RWMutex
	size     uint64
	head     *Node
	tail     *Node
	sortType SortType
}

type SortType int

const (
	_    SortType = iota
	DESC          = 0
	ASC           = 1
)

/**
 * 初始化链表
 */
func (list *List) Init(sortType SortType) {
	// 默认size
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
	(*list).mutex = new(sync.RWMutex)
	(*list).sortType = sortType
}

func (node *Node) GetData() int64 {
	if node == nil {
		return -1
	}
	return node.data
}

func (list *List) Append(data int64) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	newNode := new(Node)
	(*newNode).data = data;
	if list.size == 0 {
		// 当前没有Node 此时出入的为头结点 也同时是尾结点
		(*list).head = newNode
		(*list).tail = newNode
		(*newNode).next = nil
		(*newNode).prev = nil
		(*list).size++
	} else {
		if list.sortType == ASC {

			if data <= list.head.data {
				list.insertPrev(list.head, data)
				return
			}

			if data >= list.tail.data {
				list.insertNext(list.tail, data)
				return
			}

			// 顺序 从小到大
			for node := list.tail; node != nil; node = node.prev {
				if node.data <= data {
					list.insertNext(node, data)
					return
				}
			}
			list.insertPrev((*list).head, data)
		} else {

			if data >= list.head.data {
				list.insertPrev(list.head, data)
				return
			}

			if data <= list.tail.data {
				list.insertNext(list.tail, data)
				return
			}

			node := list.head
			if node.data <= data {
				list.insertPrev(node, data)
				return
			}
			// 顺序 从大到小
			for ; node != nil; node = node.next {
				if node.data <= data {
					list.insertPrev(node, data)
					return
				}
			}
			list.insertNext(node, data)
		}
	}
}

func (list *List) GetSize() uint64 {
	return (*list).size
}

func (list *List) GetHead() *Node {
	return (*list).head
}

func (list *List) GetTail() *Node {
	return (*list).tail
}

func (list *List) isHead(element *Node) bool {
	return list.GetHead() == element
}

func (list *List) isTail(element *Node) bool {
	return list.GetTail() == element
}

/**
 * 在节点后面插入数据
 */
func (list *List) insertNext(element *Node, data int64) bool {
	if element == nil {
		return false
	}
	if list.isTail(element) {
		newNode := new(Node)
		(*newNode).data = data
		if (*list).GetSize() == 0 {
			(*list).head = newNode
			(*list).tail = newNode
			(*newNode).prev = nil
			(*newNode).next = nil
		} else { //  挂在车队尾部
			(*newNode).prev = (*list).tail
			(*newNode).next = nil
			(*((*list).tail)).next = newNode
			(*list).tail = newNode
		}

		(*list).size++;
	} else {
		newNode := new(Node)
		(*newNode).data = data
		(*newNode).prev = element
		(*newNode).next = (*element).next
		(*element).next = newNode
		(*((*newNode).next)).prev = newNode
		(*list).size++
	}
	return true
}

/**
 * 在节点前面插入数据
 */
func (list *List) insertPrev(element *Node, data int64) bool {

	if element == nil {
		return false;
	}
	if list.isHead(element) {
		// 插入在头结点前面
		newNode := new(Node)
		(*newNode).data = data
		(*newNode).next = (*list).GetHead()
		(*newNode).prev = nil

		(*((*list).head)).prev = newNode
		(*list).head = newNode
		(*list).size ++
		return true
	} else {
		prev := (*element).prev
		return list.insertNext(prev, data)
	}
}

func (list *List) Search(data int64) *Node {
	if list.GetSize() == 0 {
		return nil
	}

	if list.sortType == ASC {
		node := new(Node)
		// 倒序 从小到大
		for node = list.tail; node != nil; {

			if node == nil {
				return nil
			}

			// 当前节点等于要查找的数据 返回当前节点
			if node.data == data {
				return node
			}
			// 当前节点不等于要查找的数据 并且当前节点是头结点 没有前置节点 返回nil
			if list.isHead(node) {
				return nil
			}

			// 要查找的数据小于头结点数据或者大于尾节点数据 则不存在改链表中（从小到大 小于最小 大于最大）
			if data < list.head.data || data > list.tail.data {
				return nil
			}

			// 要查找的数据等于头结点 返回头结点
			if data == list.head.data {
				return list.head
			}

			// 要查找的数据等于尾结点 返回尾结点
			if data == list.tail.data {
				return list.tail
			}

			//  当前节点的数据小于要查找的数据
			if node.data < data {
				//  判断前置节点的值与要查找的值
				prevNode := node.prev
				if prevNode.data > data {
					break
				} else if prevNode.data == data {
					return prevNode
				} else {
					//  前置节点的值小于要查找的值 并且前置节点不是头结点 重置节点 并接续
					if list.isHead(prevNode) {
						return nil
					}
					node = prevNode.prev
				}
			} else {
				// 节点的值小于要查找的值
				node = node.prev
			}
		}
	} else {
		node := new(Node)
		node = list.head
		if node == nil {
			return nil
		}
		if data > list.head.data || data < list.tail.data {
			return nil
		}
		if data == list.head.data {
			return list.head
		}

		if data == list.tail.data {
			return list.tail
		}

		// 倒序 从大到小
		for ; node != nil; {

			if node == nil {
				return nil
			}

			if node.data == data {
				return node
			}

			if list.isTail(node) {
				return nil
			}
			if node.data > data {
				nextNode := node.next
				if nextNode.data < data {
					return nil
				} else if nextNode.data == data {
					return nextNode
				} else {
					if list.isTail(nextNode) {
						return nil
					}
					node = nextNode.next
				}
			} else {
				node = node.next
			}
		}
	}
	return nil
}

func (list *List) Remove(element *Node) int64 {
	if element == nil {
		return 0
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	prev := (*element).prev
	next := (*element).next

	if list.isHead(element) {
		// 删除头结点
		(*list).head = next
	} else {
		(*prev).next = next;
	}

	if list.isTail(element) {
		(*list).tail = prev
	} else {
		(*next).prev = prev
	}
	(*list).size--
	return (*element).data
}

// Display 打印双链表信息
func (list *List) Display() {
	if list == nil || list.size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this list size is %d \n", list.size)
	ptr := list.head
	fmt.Printf("data is")
	for ptr != nil {
		fmt.Printf(" %v ,", ptr.data)
		ptr = ptr.next
	}
	fmt.Printf("\n")
}
