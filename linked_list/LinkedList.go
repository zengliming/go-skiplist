package linked_list

import (
	"fmt"
	"sync"
)

type Node struct {
	data uint64
	prev *Node
	next *Node
}

type List struct {
	mutex *sync.RWMutex
	size uint64
	head *Node
	tail *Node
	sortType SortType
}

type SortType int

const (
	_ SortType = iota
	DESC
	ASC
)

/**
 * 初始化链表
 */
func (list *List) Init(sortType SortType)  {
	_list := *(list)
	// 默认size
	_list.size = 0
	_list.head = nil
	_list.tail = nil
	_list.mutex = new(sync.RWMutex)
}

func (list *List) Append(data uint64)  {
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
	} else {

		if list.sortType == DESC {
			node := list.tail
			// 倒序 从小到大
			for ;node!=nil;node=node.prev {
				if node.data>data {
					list.insertNext(node, data)
				}
			}
		}else {
			node := list.head
			// 顺序 从大到小
			for ;node!=nil;node=node.next {
				if node.data>data {
					list.insertPrev(node, data)
				}
			}
		}
	}
	(*list).size++;
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
func (list *List) insertNext(element *Node, data uint64) bool {
	if element == nil {
		return false
	}

	if list.isTail(element) {
		// 插入在尾节点后面
		list.Append(data);
	}else {
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
func (list *List) insertPrev(element *Node, data uint64) bool {

	if element ==nil {
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

func (list *List) Remove(element *Node) uint64 {
	if element == nil {
		return -1
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	prev := (*element).prev
	next := (*element).next

	if list.isHead(element) {
		// 删除头结点
		(*list).head = next
	}else {
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
func (list *List)Display(){
	if list == nil || list.size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.size)
	ptr := list.head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.data)
		ptr = ptr.next
	}
}