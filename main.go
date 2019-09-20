package main

import (
	"crypto/rand"
	"fmt"
	"go-skiplist/skip_list"
	"math/big"
	"time"
)

func main() {
	list := new(skip_list.SkipList)
	list.New(2)

	var arr [10000]int64
	for i := 0; i < 50000; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		list.Append(result.Int64())
	}
	for j := 0; j < 10000; j++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		arr[j] = result.Int64()
		list.Append(result.Int64())
	}
	b1 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b1)
	for k := 0; k < 10000; k++ {
		randNum := arr[k]
		_ = list.Search(randNum)
	}
	e1 := time.Now().UnixNano()
	fmt.Printf("end %v \n", e1)
	fmt.Printf("use %v \n", (e1-b1)/10000)

	b2 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b2)
	for i := 0; i < 10000; i++ {
		randNum := arr[i]
		_ = list.DirectSearch(randNum)
	}
	e2 := time.Now().UnixNano()
	fmt.Printf("end %v \n", e2)
	fmt.Printf("use %v \n", (e2-b2)/10000)

}
