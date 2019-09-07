package skip_list

import (
	"math/rand"
	"go-skiplist/linked_list"
)

type SkipList struct {
	level int
	elements []*linked_list.List
}

func (this *SkipList) New(level int)  {
	this = new(SkipList)
	this.level = level
	for i:=0; i<level; i++ {
		list := new(linked_list.List)
		list.Init(linked_list.DESC)
		this.elements[i] = list
	}
}

func (this *SkipList) Append(data uint64)  {
	this.elements[0].Append(data)
	this.elements[rand.Intn(this.level)].Append(data)
}
