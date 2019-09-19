package skip_list

import (
	"fmt"
	"go-skiplist/linked_list"
	"math/rand"
)

type SkipList struct {
	level    int
	elements []*linked_list.List
}

func (sl *SkipList) New(level int) {
	(*sl).level = level
	for i := 0; i < level; i++ {
		list := new(linked_list.List)
		list.Init(linked_list.ASC)
		(*sl).elements = append(sl.elements, list)
	}
}

func (sl *SkipList) Append(data int64) {
	(*sl).elements[0].Append(data)
	level := rand.Intn((*sl).level-1) + 1
	(*sl).elements[level].Append(data)
}

func (sl *SkipList) Search(data int64) *linked_list.Node {

	for i := sl.level - 1; i >= 0; i-- {
		if sl.elements[i] == nil {
			continue
		}
		node := sl.elements[i].Search(data)
		if node != nil {
			return node
		}

	}
	return nil
}

func (sl *SkipList) DirectSearch(data int64) *linked_list.Node {
	node := sl.elements[0].Search(data)
	return node
}

// Display 打印双链表信息
func (sl *SkipList) Display() {
	if sl == nil {
		fmt.Println("this skip list is nil or empty")
		return
	}
	for i := 0; i < sl.level; i++ {
		(*sl).elements[i].Display()
	}
}
