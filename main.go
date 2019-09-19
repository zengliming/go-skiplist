package main

import (
	"fmt"
	"go-skiplist/skip_list"
	"math/rand"
	"time"
)

func main()  {
	list := new(skip_list.SkipList)
	list.New(3)

	rand.Seed(time.Now().UnixNano())
	for i:=0;i<1000000;i++ {
		randNum := int64(rand.Intn(1000000000))
		list.Append(randNum)
	}
	s := int64(rand.Intn(10000))
	fmt.Printf("need search data : %v \n", s)
	list.Append(s)
	b1 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b1)
	n1 := list.Search(s)
	e1 := time.Now().UnixNano()
	fmt.Printf("end %v \n", b1)
	fmt.Printf("use %v data %v \n", e1-b1, n1.GetDate())
	b2 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b2)
	n2 := list.DirectSearch(s)
	e2 := time.Now().UnixNano()
	fmt.Printf("end %v \n", e2)
	fmt.Printf("use %v  data %v \n", e2-b2, n2.GetDate())
}