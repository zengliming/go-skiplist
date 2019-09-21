package main

import (
	"crypto/rand"
	"fmt"
	"go-skiplist/skip_list"
	"math/big"
	"time"
)

func main() {
	list := skip_list.New()

	var arr [20000]int64
	for i := 0; i < 20000; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		list.Append(result.Int64())
	}
	//list.Display()
	for j := 0; j < 20000; j++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		arr[j] = result.Int64()
		list.Append(result.Int64())
	}

	b1 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b1)
	for k := 0; k < 20000; k++ {
		randNum := arr[k]
		_ = list.Search(randNum)
	}
	e1 := time.Now().UnixNano()
	fmt.Printf("end %v \n", e1)
	fmt.Printf("use %v \n", (e1-b1)/20000)

	b2 := time.Now().UnixNano()
	fmt.Printf("begin %v \n", b2)
	for i := 0; i < 20000; i++ {
		randNum := arr[i]
		_ = list.DirectSearch(randNum)
	}
	e2 := time.Now().UnixNano()
	fmt.Printf("end %v \n", e2)
	fmt.Printf("use %v \n", (e2-b2)/20000)

}
