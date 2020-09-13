package week02

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	// [-1, -1] 1
	for i := 0; i < 10000; i++ {
		res := topKFrequent([]int{-1, -1}, 1)
		if res[0] != -1 {
			fmt.Println("bad case", res)
			return
		}
	}
	fmt.Println("it's ok")
}
