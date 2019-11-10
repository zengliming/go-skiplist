package skip_list

import (
	"crypto/rand"
	"fmt"
	"go-skiplist/linked_list"
	"math/big"
)

const defaultLev = 32

type SkipList struct {
	level              int                 // level 数
	elements           []*linked_list.List // level list
	defaultRandomLevel uint                // 默认随机level
	randLev            [][]uint            //随机概率分布区间
}

// 创建skipList对象
func New() *SkipList {
	return NewWithLevel(defaultLev)
}

// 根据自定义level创建skipList对象
// @param level 自定义level 必须大于1
// @return skipList对象
func NewWithLevel(level int) *SkipList {
	sl := new(SkipList)
	lev := defaultLev
	if level > 1 {
		lev = level
	}
	sl.init(lev)
	return sl
}

// 根据level初始化skipList
// @param level level
func (sl *SkipList) init(level int) {
	(*sl).level = level
	sl.randLev = make([][]uint, level)
	for i := 0; i < level; i++ {
		// 设置list的排序规则
		list := linked_list.NewWithSort(linked_list.DESC)
		(*sl).elements = append(sl.elements, list)
		// 初始化随机概率分布区间
		sl.randLev[i] = make([]uint, 2)
		sl.randLev[i][0] = sl.sum(level - i - 1)
		sl.randLev[i][1] = sl.sum(level - i)
	}
	// 默认随机level
	(*sl).defaultRandomLevel = 1
}

// 往skipList添加数据
// @data 添加的数据
func (sl *SkipList) Append(data int64) {
	(*sl).elements[0].Append(data)
	var level uint = 0
	// 保证每一个level都有数据
	for i := sl.level - 1; i > 0; i-- {
		if sl.elements[i].GetSize() <= 0 {
			level = uint(i)
			(*sl).elements[level].Append(data)
			return
		}
	}

	// 所有level都有元素后 使用随机概率函数获取随机level
	level = sl.randomLevel()
	(*sl).elements[level].Append(data)
}

// 查找数据
// @param data 需要查找的数据
// @return *linked_list.Node 查找到的节点
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

// 直接查找数据
// @param data 需要查找的数据
// @return *linked_list.Node 查找到的节点
func (sl *SkipList) DirectSearch(data int64) *linked_list.Node {
	node := sl.elements[0].Search(data)
	return node
}

// 获取随机level 按照概率  最底下的概率越大
// @return 随机level
func (sl *SkipList) randomLevel() uint {
	// 获取l最大level
	maxLevel := sl.level
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(sl.sum(maxLevel))))
	r := uint(result.Int64())
	for i, d := range sl.randLev {
		if r >= d[0] && r < d[1] {
			if i == 0 {
				return (*sl).defaultRandomLevel
			}
			return uint(i)
		}
	}
	return (*sl).defaultRandomLevel
}

// 计算累加总和
func (sl *SkipList) sum(data int) uint {
	sum := 0
	for i := 0; i <= data; i++ {
		sum += i
	}
	return uint(sum)
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
